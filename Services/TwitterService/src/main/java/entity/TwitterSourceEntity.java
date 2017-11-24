package entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.io.Serializable;


@Entity
@Table(name = "twittersource")
public class TwitterSourceEntity implements Serializable {
    private static final long serialVersionUID = -8706689714326132798L;

    @Id
    @Column(name = "Id")
    private long id;

    @Column(name = "UserId")
    private long userId;

    @Column(name = "Name")
    private String name;

    @Column(name = "ScreenName")
    private String screenName;

    @Column(name = "Url")
    private String url;

    @Column(name = "Image")
    private String image;  // user.profile_image_url

    public TwitterSourceEntity() {
    }

    public TwitterSourceEntity(long id, long userId, String name, String screenName, String url, String image) {
        this.id = id;
        this.userId = userId;
        this.name = name;
        this.screenName = screenName;
        this.url = url;
        this.image = image;
    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public long getUserId() {
        return userId;
    }

    public void setUserId(long userId) {
        this.userId = userId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getScreenName() {
        return screenName;
    }

    public void setScreenName(String screenName) {
        this.screenName = screenName;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }
}
