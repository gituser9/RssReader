package datamodels;

import com.google.gson.JsonArray;
import java.util.List;


public class WorkData {
    private int userId;
    private JsonArray news;
    private JsonArray groups;
    private List groupIds;
    private List newsIds;

    public int getUserId() {
        return userId;
    }

    public void setUserId(int userId) {
        this.userId = userId;
    }

    public JsonArray getNews() {
        return news;
    }

    public void setNews(JsonArray news) {
        this.news = news;
    }

    public JsonArray getGroups() {
        return groups;
    }

    public void setGroups(JsonArray groups) {
        this.groups = groups;
    }

    public List getGroupIds() {
        return groupIds;
    }

    public void setGroupIds(List groupIds) {
        this.groupIds = groupIds;
    }

    public List getNewsIds() {
        return newsIds;
    }

    public void setNewsIds(List newsIds) {
        this.newsIds = newsIds;
    }
}
