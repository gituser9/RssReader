package entity;


import javax.persistence.*;
import java.io.Serializable;


@Entity
@Table(name = "twitternews")
public class TwitterNewsEntity implements Serializable {
    private static final long serialVersionUID = -8706689714326132798L;

    @Id
    @Column(name = "Id")
    private long id;

    @Column(name = "UserId")
    private long userId;

    @Column(name = "SourceId")
    private long sourceUserId;   // user.id

    @Column(name = "Text")
    private String text;

    @Column(name = "ExpandedUrl")
    private String expandedUrl;     // entities.urls[0].expanded_url

    @Column(name = "Image", nullable = true)
    private String image;

    public TwitterNewsEntity() {

    }

    public TwitterNewsEntity(long id, long userId, long sourceUserId, String text) {
        this.id = id;
        this.userId = userId;
        this.sourceUserId = sourceUserId;
        this.text = text;
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

    public long getSourceUserId() {
        return sourceUserId;
    }

    public void setSourceUserId(long sourceUserId) {
        this.sourceUserId = sourceUserId;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public String getExpandedUrl() {
        return expandedUrl;
    }

    public void setExpandedUrl(String expandedUrl) {
        this.expandedUrl = expandedUrl;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }
}
