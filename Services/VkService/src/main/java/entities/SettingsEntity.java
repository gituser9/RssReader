package entities;

import javax.persistence.*;
import java.io.Serializable;


@Entity
@Table(name = "settings")
public class SettingsEntity implements Serializable {
    private static final long serialVersionUID = -8706689714326132798L;

    @Id
    @Column(name = "Id")
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;

    @Column(name = "UserId")
    private long userId;

    @Column(name = "VkNewsEnabled")
    private boolean vkNewsEnabled;

    @SuppressWarnings("UnusedDeclaration")
    public SettingsEntity() {
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

    public boolean isVkNewsEnabled() {
        return vkNewsEnabled;
    }

    public void setVkNewsEnabled(boolean vkNewsEnabled) {
        this.vkNewsEnabled = vkNewsEnabled;
    }
}
