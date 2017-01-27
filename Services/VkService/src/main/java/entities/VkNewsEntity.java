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
    private int userId;

    @Column(name = "GroupId")
    private int groupId;

    @Column(name = "PostId")
    private int postId;

    @Column(name = "Text", updatable = false)
    private String text;

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

    public int getUserId() {
        return userId;
    }

    public void setUserId(int userId) {
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

    @Override
    public String toString() {
        return "VkNewsEntity{" +
                "id=" + id +
                ", userId=" + userId +
                ", text='" + text + '\'' +
                '}';
    }
}
