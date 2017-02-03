import datamodels.AppProperties;
import datamodels.WorkData;
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
        Properties prop = new Properties();
        AppProperties appProperties = new AppProperties();

        try (InputStream input = new FileInputStream("config.properties")) {
            // load a properties file
            prop.load(input);

            appProperties.setClientId(prop.getProperty("clientId"));
            appProperties.setClientSecret(prop.getProperty("clientSecret"));
            appProperties.setPasswordSalt(prop.getProperty("passwordSalt"));
        } catch (IOException e) {
            System.out.println("Load properties error: " + e.getMessage());
        }

        VkService vkService = new VkService();
        NetService netService = new NetService(appProperties);

        try (Session session = HibernateSessionFactory.getSessionFactory().openSession()) {
            Criteria criteria = session.createCriteria(UserEntity.class);
            criteria.add(Restrictions.eq("vkNewsEnabled", true));

            List<UserEntity> users = (List<UserEntity>) criteria.list();

            for (UserEntity user : users) {
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
                WorkData workData = netService.getNews(user.getVkLogin(), user.getVkPassword());
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
}
