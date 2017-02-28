package services;

import com.google.gson.JsonArray;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParser;
import datamodels.AppProperties;
import datamodels.WorkData;
import org.apache.http.HttpEntity;
import org.apache.http.HttpResponse;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.impl.client.HttpClientBuilder;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;


public class NetService {

    private HttpClient httpClient = HttpClientBuilder.create().build();
    private HttpResponse response;
    private HttpEntity entity;
    private JsonParser parser = new JsonParser();
    private AppProperties appProperties;

    public NetService(AppProperties appProperties) {
        this.appProperties = appProperties;
    }

    public WorkData getNews(String login, String password) {
        String token = getToken(login, password);
        HttpGet get = new HttpGet("https://api.vk.com/method/newsfeed.get?filters=post&access_token=" + token);

        try {
            response = httpClient.execute(get);
        } catch (IOException e) {
            e.printStackTrace();
        }

        entity = response.getEntity();
        StringBuilder builder = new StringBuilder();
        String line;

        try (BufferedReader buffer = new BufferedReader(new InputStreamReader(entity.getContent()))) {
            while ((line = buffer.readLine()) != null) {
                builder.append(line);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        JsonObject newsJson = parser.parse(builder.toString()).getAsJsonObject();
        JsonObject responseJson = newsJson.get("response").getAsJsonObject();

        WorkData workData = new WorkData();
        workData.setNews(responseJson.get("items").getAsJsonArray());
        workData.setGroups(responseJson.get("groups").getAsJsonArray());

        return workData;
    }

    private String getToken(String login, String password) {
        String url  = String.format(
                "https://oauth.vk.com/token?grant_type=password&client_id=%s&client_secret=%s&username=%s&password=%s",
                appProperties.getClientId(), appProperties.getClientSecret(), login, password
            );
        HttpPost post = new HttpPost(url);

        try {
            response = httpClient.execute(post);
        } catch (IOException e) {
            e.printStackTrace();
        }

        entity = response.getEntity();
        String line;
        StringBuilder builder = new StringBuilder();

        try (BufferedReader buffer = new BufferedReader(new InputStreamReader(entity.getContent()))) {
            while ((line = buffer.readLine()) != null) {
                builder.append(line);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        JsonObject accessJson = parser.parse(builder.toString()).getAsJsonObject();

        return accessJson.get("access_token").getAsString();
    }
}
