import com.google.gson.*;
import model.AppProperties;
import model.TwitterNews;
import okhttp3.*;
import org.codehaus.plexus.util.Base64;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;


public class DataProvider {

    private String token;
    private String tokenType;
    private final String contentType = "application/x-www-form-urlencoded;charset=UTF-8";
    private final MediaType mediaType = MediaType.parse(contentType);
    private final Gson gson = new Gson();
    private final OkHttpClient client = new OkHttpClient();
    private final AppProperties appProperties;

    public DataProvider(AppProperties appProperties) {
        this.appProperties = appProperties;
    }

    /**
     * Get user tweets
     * @param screenName String
     * @return
     * @throws IOException
     */
    public JsonArray getNewsForUser(String screenName) throws IOException {
        List<String> friendScreenNames = getFriendScreenNames(screenName);
        String urlTemplate = "https://api.twitter.com/1.1/statuses/user_timeline.json?include_rts=1&exclude_replies=1&screen_name=";
        String authString = tokenType + " " + token;
        Gson gson = new Gson();
        JsonArray allNews = new JsonArray();

        for (String friendScreenName : friendScreenNames) {
            Request request = new Request.Builder()
                    .header("Authorization", authString)
                    .url(urlTemplate + friendScreenName)
                    .build();
            Response response = client.newCall(request).execute();
            JsonArray newsArray = gson.fromJson(response.body().string(), JsonArray.class);
            allNews.addAll(newsArray);
        }

        return allNews;
    }

    /**
     * Get app token
     * @throws IOException
     */
    public void getToken() throws IOException {
//        String authString = "7ShgcCaLcxFebeHNxlCvWuzRs:rDteEJWGVUZLiSuMHYtrzls9rhOpHSpyPLJ88Mwc9BW400J9lk";
        String authString = appProperties.getClientId() + ":" + appProperties.getClientSecret();
        byte[] base64 = Base64.encodeBase64(authString.getBytes());
        RequestBody requestBody = RequestBody.create(mediaType, "grant_type=client_credentials");
        String tokenUrl = "https://api.twitter.com/oauth2/token";
        Request request = new Request.Builder()
            .header("Authorization", "Basic " + new String(base64))
            .header("Content-Type", contentType)
            .url(tokenUrl)
            .post(requestBody)
            .build();
        Response response = client.newCall(request).execute();
        JsonObject jsonObject = gson.fromJson(response.body().string(), JsonObject.class);
        System.out.println(jsonObject);
        token = jsonObject.get("access_token").getAsString();
        tokenType = jsonObject.get("token_type").getAsString();
    }

    /**
     * Get screen names user friends
     * @param screenName String
     * @return List<String>
     * @throws IOException
     */
    private List<String> getFriendScreenNames(String screenName) throws IOException {
        String url = "https://api.twitter.com/1.1/friends/list.json?cursor=-1&screen_name="+screenName+"&skip_status=true&include_user_entities=false";
        String authString = tokenType + " " + token;
        Request request = new Request.Builder()
                .header("Authorization", authString)
                .url(url)
                .build();
        Response response = client.newCall(request).execute();
        JsonObject json = gson.fromJson(response.body().string(), JsonObject.class);
        JsonArray users = json.get("users").getAsJsonArray();
        List<String> friendScreenNames = new ArrayList<String>();

        for (JsonElement user : users) {
            JsonObject object = user.getAsJsonObject();
            friendScreenNames.add(object.get("screen_name").getAsString());
        }

        return friendScreenNames;
    }

}
