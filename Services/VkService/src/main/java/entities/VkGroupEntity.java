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
    private int userId;

    public VkGroupEntity() {

    }

    public VkGroupEntity(int gid, int userId, String name, String linkedName) {
        this.gid = gid;
        this.userId = userId;
        this.name = name;
        this.linkedName = linkedName;
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

    public int getUserId() {
        return userId;
    }

    public void setUserId(int userId) {
        this.userId = userId;
    }
}
