import com.google.gson.JsonArray;
import entity.SettingsEntity;
import entity.UserEntity;
import model.AppProperties;
import org.hibernate.Criteria;
import org.hibernate.Session;
import org.hibernate.criterion.Projections;
import org.hibernate.criterion.Restrictions;
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
            System.out.println("No properties");
            return;
        }

        do {
            try {
                appProperties = getProperties();

                if (appProperties == null) {
                    continue;
                }

                updateTwitterNews(appProperties);
                Thread.sleep(appProperties.getSleepMinutes() * 1000 * 60);
            } catch (Exception e) {
                e.printStackTrace();
            }
        } while (true);
    }

    public static void updateTwitterNews(AppProperties appProperties) {
        DataProvider dataProvider = new DataProvider(appProperties);
        HibernateSessionFactory.buildSessionFactory(appProperties);
        Session session = HibernateSessionFactory.getSessionFactory().openSession();
        Criteria settingsCriteria = session.createCriteria(SettingsEntity.class);
        settingsCriteria.add(Restrictions.eq("twitterNewsEnabled", true));
        settingsCriteria.setProjection(Projections.property("userId"));
        Object[] userIds = settingsCriteria.list().toArray();

        if (userIds.length == 0) {
            return;
        }

        Criteria usersCriteria = session.createCriteria(UserEntity.class);
        usersCriteria.add(Restrictions.in("id", userIds));
        List<UserEntity> users = (List<UserEntity>) usersCriteria.list();

        TwitterService twitterService = new TwitterService(session);

        try {
            dataProvider.getToken();

            for (UserEntity user : users) {
                JsonArray json = dataProvider.getNewsForUser(user.getTwitterScreenName());
                twitterService.update(json, user.getId());
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            session.close();
            HibernateSessionFactory.shutdown();
        }
    }

    private static AppProperties getProperties() {
        Properties properties = new Properties();
        AppProperties appProperties = new AppProperties();

        try (InputStream input = new FileInputStream("twitterconfig.properties")) {
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
