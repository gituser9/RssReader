import com.google.gson.JsonArray;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.sun.istack.internal.Nullable;
import entity.SettingsEntity;
import entity.TwitterNewsEntity;
import entity.TwitterSourceEntity;
import org.hibernate.Criteria;
import org.hibernate.SQLQuery;
import org.hibernate.Session;
import org.hibernate.criterion.Projections;
import org.hibernate.criterion.Restrictions;
import utils.HibernateSessionFactory;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class TwitterService {
    private final Session session;
    private List<Long> existingNewsIds;
    private List<Long> existingSourceIds;
    private long userId;

    public TwitterService(Session session) {
        this.session = session;
    }

    public void update(JsonArray json, long userId) {
        this.userId = userId;
        prepareData();
        saveNews(json);
    }

    private List<TwitterNewsEntity> getNews(JsonArray json) {
        List<TwitterNewsEntity> newsEntities = new ArrayList<>(json.size());

        for (JsonElement item : json) {
            JsonObject newsObject = item.getAsJsonObject();
            Long id = newsObject.get("id").getAsLong();

            if (Collections.binarySearch(existingNewsIds, id) >= 0) {
                continue;
            }

            long sourceId = newsObject.get("user").getAsJsonObject().get("id").getAsLong();
            String text = newsObject.get("text").getAsString();
            TwitterNewsEntity entity = new TwitterNewsEntity(id, userId, sourceId, text);
            entity.setExpandedUrl(getExpandedUrl(newsObject));

            newsEntities.add(entity);
        }

        return newsEntities;
    }

    private List<TwitterSourceEntity> getSources(JsonArray json) {
        List<TwitterSourceEntity> sourcesEntities = new ArrayList<>();

        for (JsonElement item : json) {
            JsonObject newsObject = item.getAsJsonObject();
            long sourceId = newsObject.get("user").getAsJsonObject().get("id").getAsLong();

            if (existingSourceIds.contains(sourceId)) {
                continue;
            }

            JsonObject userObject = newsObject.get("user").getAsJsonObject();
            String name = userObject.get("name").getAsString();
            String screenName = userObject.get("screen_name").getAsString();
            String url = userObject.get("url").getAsString();
            String image = userObject.get("profile_image_url").getAsString();
            TwitterSourceEntity entity = new TwitterSourceEntity(sourceId, userId, name, screenName, url, image);

            sourcesEntities.add(entity);
        }

        return sourcesEntities;
    }

    @Nullable
    private String getExpandedUrl(JsonObject json) {
        if (!json.has("entities")) {
            return null;
        }

        JsonObject entitiesObject = json.get("entities").getAsJsonObject();

        if (!entitiesObject.has("urls")) {
            return null;
        }

        JsonArray urls = entitiesObject.get("urls").getAsJsonArray();
        JsonObject url = urls.get(0).getAsJsonObject();

        return url.get("expanded_url").getAsString();
    }

    private void prepareData() {
        Criteria newsCriteria = session.createCriteria(TwitterNewsEntity.class);
        newsCriteria.add(Restrictions.eq("userId", userId));
        newsCriteria.setProjection(Projections.property("id"));

        Criteria sourceCriteria = session.createCriteria(TwitterSourceEntity.class);
        newsCriteria.add(Restrictions.eq("userId", userId));
        newsCriteria.setProjection(Projections.property("id"));

        existingNewsIds = newsCriteria.list();
        existingSourceIds = sourceCriteria.list();
        Collections.sort(existingNewsIds);
    }

    private void saveNews(JsonArray json) {
        // save news
        session.beginTransaction();
        SQLQuery query = session.createSQLQuery("SET NAMES utf8mb4");
        query.executeUpdate();

        for (TwitterNewsEntity item : getNews(json)) {
            try {
                session.save(item);
            } catch (Exception e) {
                e.printStackTrace();
            }
        }

        session.getTransaction().commit();

        // save groups
        session.beginTransaction();

        for (TwitterSourceEntity item : getSources(json)) {
            try {
                session.save(item);
            } catch (Exception e) {
                e.printStackTrace();
            }
        }

        session.getTransaction().commit();
    }
}
