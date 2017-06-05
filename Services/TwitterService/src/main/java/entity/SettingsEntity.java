package entity;

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

    @Column(name = "RssEnabled")
    private boolean rssEnabled;

    @Column(name = "UnreadOnly")
    private boolean unreadOnly;

    @Column(name = "MarkSameRead")
    private boolean markSameRead;

    @Column(name = "TwitterEnabled")
    private boolean twitterNewsEnabled;

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

    public boolean isRssEnabled() {
        return rssEnabled;
    }

    public void setRssEnabled(boolean rssEnabled) {
        this.rssEnabled = rssEnabled;
    }

    public boolean isUnreadOnly() {
        return unreadOnly;
    }

    public void setUnreadOnly(boolean unreadOnly) {
        this.unreadOnly = unreadOnly;
    }

    public boolean isMarkSameRead() {
        return markSameRead;
    }

    public void setMarkSameRead(boolean markSameRead) {
        this.markSameRead = markSameRead;
    }

    public boolean isTwitterNewsEnabled() {
        return twitterNewsEnabled;
    }

    public void setTwitterNewsEnabled(boolean twitterNewsEnabled) {
        this.twitterNewsEnabled = twitterNewsEnabled;
    }
}
