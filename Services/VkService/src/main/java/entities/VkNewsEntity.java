package entities;

import javax.persistence.*;
import java.io.Serializable;


@Entity
@Table(name = "vknews")
public class VkNewsEntity implements Serializable {
    private static final long serialVersionUID = -8706689714326132798L;

    @Id
    @Column(name = "Id")
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;

    @Column(name = "UserId")
    private long userId;

    @Column(name = "GroupId")
    private int groupId;

    @Column(name = "PostId")
    private int postId;

    @Column(name = "Text", updatable = false)
    private String text;

    @Column(name = "Image", updatable = false)
    private String image;

    @Column(name = "Link", updatable = false)
    private String link;

    @Column(name = "Timestamp", updatable = false)
    private long timestamp;

    @SuppressWarnings("UnusedDeclaration")
    public VkNewsEntity() {
    }

    @SuppressWarnings("UnusedDeclaration")
    public VkNewsEntity(String text) {
        setText(text);
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public long getUserId() {
        return userId;
    }

    public void setUserId(long userId) {
        this.userId = userId;
    }

    public int getGroupId() {
        return groupId;
    }

    public void setGroupId(int groupId) {
        this.groupId = groupId;
    }

    public int getPostId() {
        return postId;
    }

    public void setPostId(int postId) {
        this.postId = postId;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }

    public String getLink() {
        return link;
    }

    public void setLink(String link) {
        this.link = link;
    }

    public long getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(long timestamp) {
        this.timestamp = timestamp;
    }

    @Override
    public String toString() {
        return "VkNewsEntity{" +
                "id=" + id +
                ", userId=" + userId +
                ", text='" + text + '\'' +
                '}';
    }
}
