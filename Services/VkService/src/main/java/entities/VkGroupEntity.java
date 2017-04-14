package entities;

import javax.persistence.*;
import java.io.Serializable;


@Entity
@Table(name = "vkgroups")
public class VkGroupEntity implements Serializable {
    private static final long serialVersionUID = -8706689714326132798L;

    @Id
    @Column(name = "Id")
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;

    @Column(name = "Gid")
    private int gid;

    @Column(name = "Name", unique = false, updatable = true)
    private String name;

    @Column(name = "LinkedName", unique = false, updatable = true)
    private String linkedName;

    @Column(name = "UserId")
    private long userId;

    @Column(name = "Image", unique = false, updatable = true)
    private String image;

    public VkGroupEntity() {

    }

    public VkGroupEntity(int gid, long userId, String name, String linkedName, String image) {
        this.gid = gid;
        this.userId = userId;
        this.name = name;
        this.linkedName = linkedName;
        this.image = image;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getGid() {
        return gid;
    }

    public void setGid(int gid) {
        this.gid = gid;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getLinkedName() {
        return linkedName;
    }

    public void setLinkedName(String linkedName) {
        this.linkedName = linkedName;
    }

    public long getUserId() {
        return userId;
    }

    public void setUserId(long userId) {
        this.userId = userId;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }
}
