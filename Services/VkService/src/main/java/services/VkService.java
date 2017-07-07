package services;

import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.sun.istack.NotNull;
import datamodels.AppProperties;
import datamodels.ImageLink;
import datamodels.WorkData;
import entities.VkGroupEntity;
import entities.VkNewsEntity;
import org.hibernate.Criteria;
import org.hibernate.SQLQuery;
import org.hibernate.Session;
import org.hibernate.criterion.Restrictions;
import utils.HibernateSessionFactory;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;


public class VkService {

    public VkService(AppProperties appProperties) {
        AppProperties appProperties1 = appProperties;
    }

    public void saveNews(WorkData workData, Session session) {
        List<VkNewsEntity> data = convertData(workData, session);
        saveData(data, session);
    }

    private List<VkNewsEntity> convertData(WorkData workData, Session session) {
        List<Integer> groupIds = (List<Integer>) workData.getGroupIds();
        List<Integer> newsIds = (List<Integer>) workData.getNewsIds();
        Collections.sort(newsIds);

        List<VkNewsEntity> result = new ArrayList<>(workData.getNews().size());
        List<Integer> updateGroups = new ArrayList<>(groupIds.size());
        long userId = workData.getUserId();

        // get new only and convert
        for (JsonElement item : workData.getNews()) {
            JsonObject json = item.getAsJsonObject();

            if (json.get("marked_as_ads").getAsInt() == 1) {
                continue;
            }

            ImageLink imageLink = getImageLink(json);
            Integer postId = json.get("post_id").getAsInt();
            Integer groupId = -json.get("source_id").getAsInt();

            if (Collections.binarySearch(newsIds, postId) >= 0) {
                continue;
            }

            if (!groupIds.contains(groupId)) {
                // add new group for user
                addGroup(groupId, workData, session);
                groupIds.add(groupId);
            } else {
                updateGroups.add(groupId);
            }

            VkNewsEntity entity = new VkNewsEntity();
            entity.setGroupId(groupId);
            entity.setPostId(postId);
            entity.setText(json.get("text").getAsString());
            entity.setUserId(userId);
            entity.setImage(imageLink.getImage());
            entity.setLink(imageLink.getLink());
            entity.setTimestamp(json.get("date").getAsLong());

            result.add(entity);
        }

        //cleanGroups(workData, groupIds, session);  // todo: in thread
        updateGroups(workData, groupIds, session);

        return result;
    }

    private void updateGroups(WorkData workData, List<Integer> groupIds, Session session) {
        session.beginTransaction();
        Criteria groupCriteria = session.createCriteria(VkGroupEntity.class);
        groupCriteria.add(Restrictions.in("id", groupIds));
        List<VkGroupEntity> entities = groupCriteria.list();

        for (VkGroupEntity entity : entities) {
            for (JsonElement item : workData.getGroups()) {
                JsonObject json = item.getAsJsonObject();
                int id = json.get("gid").getAsInt();

                if (entity.getGid() == id) {
                    entity.setName(json.get("name").getAsString());
                    entity.setLinkedName(json.get("screen_name").getAsString());
                    entity.setImage(json.get("photo").getAsString());

                    try {
                        session.save(entity);
                    } catch (Exception e) {
                        continue;
                    }
                }
            }
        }

        try {
            session.getTransaction().commit();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private void saveData(List<VkNewsEntity> news, Session session) {
        session.beginTransaction();
        SQLQuery query = session.createSQLQuery("SET NAMES utf8mb4");
        query.executeUpdate();

        for (VkNewsEntity item : news) {
            try {
                session.save(item);
            } catch (Exception e) {
                continue;
            }
        }

        try {
            session.getTransaction().commit();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private void addGroup(int groupId, WorkData workData, Session session) {
        for (JsonElement item : workData.getGroups()) {
            JsonObject json = item.getAsJsonObject();
            int id = json.get("gid").getAsInt();

            if (groupId == id) {
                String name = json.get("name").getAsString();
                String linkedName = json.get("screen_name").getAsString();
                String image = json.get("photo").getAsString();

                session.beginTransaction();
                VkGroupEntity groupEntity = new VkGroupEntity(groupId, workData.getUserId(), name, linkedName, image);
                session.save(groupEntity);
                session.getTransaction().commit();

                return;
            }
        }
    }

    // delete group from DB if any group not in JSON
    private void cleanGroups(WorkData workData, List<Integer> groupIds, Session session) {

        // get all grous for user
        Criteria criteria = session.createCriteria(VkGroupEntity.class);
        criteria.add(Restrictions.eq("userId", workData.getUserId()));

        List<VkGroupEntity> groups = (List<VkGroupEntity>) criteria.list();
        List<VkGroupEntity> oldGroups = new ArrayList<>();

        // get old groups
        for (VkGroupEntity group : groups) {
            if (!groupIds.contains(group.getGid())) {
                oldGroups.add(group);
            }
        }

        if (oldGroups.size() == 0) {
            return;
        }

        // delete all old groups if it is not in groupIds
        session.beginTransaction();

        for (VkGroupEntity vkGroupEntity : oldGroups) {
            session.delete(vkGroupEntity);
        }

        session.getTransaction().commit();
    }

    @NotNull
    private ImageLink getImageLink(JsonObject json) {
        if (!json.has("attachment")) {
            return new ImageLink("", "");
        }


        JsonObject attachment = json.get("attachment").getAsJsonObject();
        String postType = attachment.get("type").getAsString();
        String image = null;
        String link = null;

        if (postType.equals("photo")) {
            if (attachment.has("photo")) {
                JsonObject photo = attachment.getAsJsonObject("photo");

                if (photo.has("src_big")){
                    image = photo.get("src_big").getAsString();
                }
            }
        } else if (postType.equals("link")) {
            JsonObject linkObject = attachment.getAsJsonObject("link");

            if (linkObject.has("image_big")) {
                image = linkObject.get("image_big").getAsString();
            } else if (linkObject.has("image_src")) {
                image = linkObject.get("image_src").getAsString();
            }

            link = linkObject.get("url").getAsString();
        } else if (postType.equals("doc")) {
            JsonObject doc = attachment.get("doc").getAsJsonObject();

            if (doc.has("ext") && doc.get("ext").getAsString().equals("gif")) {
                image = doc.get("url").getAsString();
            }
        }

        return new ImageLink(image, link);
    }

}
