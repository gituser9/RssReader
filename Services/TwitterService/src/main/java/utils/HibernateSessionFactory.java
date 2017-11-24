package utils;


import entity.SettingsEntity;
import entity.TwitterNewsEntity;
import entity.TwitterSourceEntity;
import entity.UserEntity;
import model.AppProperties;
import org.hibernate.SessionFactory;
import org.hibernate.boot.registry.StandardServiceRegistryBuilder;
import org.hibernate.cfg.Configuration;
import org.hibernate.service.ServiceRegistry;


public class HibernateSessionFactory {
    private static SessionFactory sessionFactory;

    public static void buildSessionFactory(AppProperties appProperties) {
//        return sessionFactory;
        Configuration configuration = confugurationBuilder(appProperties);
        StandardServiceRegistryBuilder builder = new StandardServiceRegistryBuilder();
        builder.applySettings(configuration.getProperties());
        ServiceRegistry serviceRegistry = builder.build();

        sessionFactory = configuration.buildSessionFactory(serviceRegistry);
    }

    private static Configuration confugurationBuilder(AppProperties appProperties) {
        switch (appProperties.getDbEngine()) {
            case "mysql":
                return createMysqlConfiguration(appProperties);
            default:
                return null;
        }

    }

    public static SessionFactory getSessionFactory() {
        return sessionFactory;
    }

    private static Configuration createMysqlConfiguration(AppProperties appProperties) {
        Configuration configuration = new Configuration();
        configuration.addAnnotatedClass(UserEntity.class);
        configuration.addAnnotatedClass(SettingsEntity.class);
        configuration.addAnnotatedClass(TwitterNewsEntity.class);
        configuration.addAnnotatedClass(TwitterSourceEntity.class);

        configuration.setProperty("hibernate.dialect", "org.hibernate.dialect.MySQLDialect");
        configuration.setProperty("hibernate.connection.driver_class", "com.mysql.jdbc.Driver");
        configuration.setProperty("hibernate.connection.url", appProperties.getHibernateConnectionString());
        configuration.setProperty("hibernate.connection.username", appProperties.getDbLogin());
        configuration.setProperty("hibernate.connection.password", appProperties.getDbPassword());

        configuration.setProperty("hibernate.connection.CharSet", "utf8");
        configuration.setProperty("hibernate.connection.characterEncoding", "utf8");
        configuration.setProperty("hibernate.connection.useUnicode", "true");

        configuration.setProperty("hibernate.show_sql", "false");


        return configuration;
    }

    public static void shutdown() {
        // Close caches and connection pools
        getSessionFactory().close();
    }


}
