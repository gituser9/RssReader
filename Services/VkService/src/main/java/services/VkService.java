package services;

import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import datamodels.WorkData;
import entities.VkGroupEntity;
import entities.VkNewsEntity;
import org.hibernate.SQLQuery;
import org.hibernate.Session;
import utils.HibernateSessionFactory;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;


public class VkService {

    public void saveNews(WorkData workData, Session session) {
        List<VkNewsEntity> data = convertData(workData, session);
        saveData(data);
    }

    private List<VkNewsEntity> convertData(WorkData workData, Session session) {
        List<Integer> groupIds = (List<Integer>) workData.getGroupIds();
        List<Integer> newsIds = (List<Integer>) workData.getNewsIds();
        Collections.sort(groupIds);
        Collections.sort(newsIds);

        List<VkNewsEntity> result = new ArrayList<>(workData.getNews().size());
        int userId = workData.getUserId();

        // get new only and convert
        for (JsonElement item : workData.getNews()) {
            JsonObject json = item.getAsJsonObject();
            Integer postId = json.get("post_id").getAsInt();
            Integer groupId = -json.get("source_id").getAsInt();

            if (Collections.binarySearch(newsIds, postId) >= 0) {
                continue;
            }

            if (Collections.binarySearch(groupIds, groupId) == -1) {
                // add new group for user
                addGroup(groupId, workData, session);
                groupIds.add(groupId);
                Collections.sort(groupIds);
            }

            VkNewsEntity entity = new VkNewsEntity();
            entity.setGroupId(groupId);
            entity.setPostId(postId);
            entity.setText(json.get("text").getAsString());
            entity.setUserId(userId);

            result.add(entity);
        }

        cleanGroups(workData, groupIds, session);  // todo: in thread

        return result;
    }

    private void saveData(List<VkNewsEntity> news) {
        try (Session session = HibernateSessionFactory.getSessionFactory().openSession()) {
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

                session.beginTransaction();
                VkGroupEntity groupEntity = new VkGroupEntity(groupId, workData.getUserId(), name, linkedName);
                session.save(groupEntity);
                session.getTransaction().commit();

                return;
            }
        }
    }

    // delete group from DB if any group not in JSON
    private void cleanGroups(WorkData workData, List<Integer> groupIds, Session session) {

        // get all grous for user

        // get old groups

        // delete all old groups if it is not in groupIds

    }

}