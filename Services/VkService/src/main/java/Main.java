import com.sun.istack.internal.Nullable;
import datamodels.AppProperties;
import datamodels.WorkData;
import entities.SettingsEntity;
import entities.UserEntity;
import entities.VkGroupEntity;
import entities.VkNewsEntity;
import org.hibernate.Criteria;
import org.hibernate.Session;
import org.hibernate.criterion.Projections;
import org.hibernate.criterion.Restrictions;
import services.NetService;
import services.VkService;
import utils.HibernateSessionFactory;

import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.List;
import java.util.Properties;


public class Main {

    public static void main(String[] args) {
        AppProperties appProperties = getProperties();

        if (appProperties == null) {
            return;
        }

        updateVkNews(appProperties);

        while (true) {
            try {
                appProperties = getProperties();

                if (appProperties == null) {
                    continue;
                }

                Thread.sleep(appProperties.getSleepMinutes() * 1000 * 60);
                updateVkNews(appProperties);
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }

    private static void updateVkNews(AppProperties appProperties) {
        System.out.println("START");

        VkService vkService = new VkService(appProperties);
        NetService netService = new NetService(appProperties);

        try (Session session = HibernateSessionFactory.getSessionFactory(appProperties).openSession()) {
            Criteria settingsCriteria = session.createCriteria(SettingsEntity.class);
            settingsCriteria.add(Restrictions.eq("vkNewsEnabled", true));
            settingsCriteria.setProjection(Projections.property("userId"));
            Object[] userIds = settingsCriteria.list().toArray();

            Criteria usersCriteria = session.createCriteria(UserEntity.class);
            usersCriteria.add(Restrictions.in("id", userIds));

            List<UserEntity> users = (List<UserEntity>) usersCriteria.list();

            for (UserEntity user : users) {
                String vkPassword = user.getVkPassword();

                if (vkPassword == null) {
                    continue;
                }

                // query criteria
                Criteria groupCriteria = session.createCriteria(VkGroupEntity.class);
                Criteria newsCriteria = session.createCriteria(VkNewsEntity.class);
                groupCriteria.add(Restrictions.eq("userId", user.getId()));
                groupCriteria.setProjection(Projections.property("gid"));
                newsCriteria.add(Restrictions.eq("userId", user.getId()));
                newsCriteria.setProjection(Projections.property("postId"));

                // lists of existing groups and news objects
                List existingGroupList = groupCriteria.list();
                List existingNewsList = newsCriteria.list();

                // get new news
                WorkData workData = netService.getNews(user.getVkLogin(), vkPassword);
                workData.setGroupIds(existingGroupList);
                workData.setNewsIds(existingNewsList);
                workData.setUserId(user.getId());
                vkService.saveNews(workData, session);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        System.out.println("END");
    }

    @Nullable
    private static AppProperties getProperties() {
        Properties properties = new Properties();
        AppProperties appProperties = new AppProperties();

        try (InputStream input = new FileInputStream("vkconfig.properties")) {
            // load a properties file
            properties.load(input);

            appProperties.setClientId(properties.getProperty("clientId"));
            appProperties.setClientSecret(properties.getProperty("clientSecret"));
            appProperties.setSleepMinutes(Integer.parseInt(properties.getProperty("sleepMinutes")));
            appProperties.setDbEngine(properties.getProperty("dbEngine"));
            appProperties.setDbLogin(properties.getProperty("dbLogin"));
            appProperties.setDbPassword(properties.getProperty("dbPassword"));
            appProperties.setHibernateConnectionString(properties.getProperty("hibernateConnectionString"));
        } catch (IOException e) {
            System.out.println("Load properties error: " + e.getMessage());
            return null;
        }

        return appProperties;
    }
}
