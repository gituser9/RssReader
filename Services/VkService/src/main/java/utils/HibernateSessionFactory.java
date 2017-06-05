package utils;

import datamodels.AppProperties;
import entities.SettingsEntity;
import entities.UserEntity;
import entities.VkGroupEntity;
import entities.VkNewsEntity;
import org.hibernate.SessionFactory;
import org.hibernate.boot.registry.StandardServiceRegistryBuilder;
import org.hibernate.cfg.Configuration;
import org.hibernate.service.ServiceRegistry;


public class HibernateSessionFactory {

    private static SessionFactory sessionFactory;

    public static SessionFactory getSessionFactory(AppProperties appProperties) {
//        return sessionFactory;
        Configuration configuration = confugurationBuilder(appProperties);
        StandardServiceRegistryBuilder builder = new StandardServiceRegistryBuilder();
        builder.applySettings(configuration.getProperties());
        ServiceRegistry serviceRegistry = builder.build();

        sessionFactory = configuration.buildSessionFactory(serviceRegistry);

        return sessionFactory;
    }

    public static void shutdown() {
        // Close caches and connection pools
        if (sessionFactory != null) {
            sessionFactory.close();
        }
    }

    private static Configuration confugurationBuilder(AppProperties appProperties) {
        switch (appProperties.getDbEngine()) {
            case "mysql":
                return createMysqlConfiguration(appProperties);
            default:
                return null;
        }

    }

    private static Configuration createMysqlConfiguration(AppProperties appProperties) {
        Configuration configuration = new Configuration();
        configuration.addAnnotatedClass(VkNewsEntity.class);
        configuration.addAnnotatedClass(VkGroupEntity.class);
        configuration.addAnnotatedClass(SettingsEntity.class);
        configuration.addAnnotatedClass(UserEntity.class);

        configuration.setProperty("hibernate.dialect", "org.hibernate.dialect.MySQLDialect");
        configuration.setProperty("hibernate.connection.driver_class", "com.mysql.jdbc.Driver");
        configuration.setProperty("hibernate.connection.url", appProperties.getHibernateConnectionString());
        configuration.setProperty("hibernate.connection.username", appProperties.getDbLogin());
        configuration.setProperty("hibernate.connection.password", appProperties.getDbPassword());
        configuration.setProperty("hibernate.show_sql", "false");


        return configuration;
    }


}
