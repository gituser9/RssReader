package com.newshub.newshub_android.settings.model;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.io.Serializable;

public class Settings implements Serializable {
    @SerializedName("UserId")
    @Expose
    private int userId;
    @SerializedName("VkLogin")
    @Expose
    private String vkLogin;
    @SerializedName("VkPassword")
    @Expose
    private String vkPassword;
    @SerializedName("TwitterName")
    @Expose
    private String twitterName;
    @SerializedName("VkNewsEnabled")
    @Expose
    private boolean vkNewsEnabled;
    @SerializedName("TwitterEnabled")
    @Expose
    private boolean twitterEnabled;
    @SerializedName("MarkSameRead")
    @Expose
    private boolean markSameRead;
    @SerializedName("UnreadOnly")
    @Expose
    private boolean unreadOnly;
    @SerializedName("RssEnabled")
    @Expose
    private boolean rssEnabled;
    @SerializedName("ShowPreviewButton")
    @Expose
    private boolean showPreviewButton;
    @SerializedName("ShowTabButton")
    @Expose
    private boolean showTabButton;
    @SerializedName("ShowReadButton")
    @Expose
    private boolean showReadButton;
    private final static long serialVersionUID = -5007899797193450585L;

    public int getUserId() {
        return userId;
    }

    public void setUserId(int userId) {
        this.userId = userId;
    }

    public String getVkLogin() {
        return vkLogin;
    }

    public void setVkLogin(String vkLogin) {
        this.vkLogin = vkLogin;
    }

    public String getVkPassword() {
        return vkPassword;
    }

    public void setVkPassword(String vkPassword) {
        this.vkPassword = vkPassword;
    }

    public String getTwitterName() {
        return twitterName;
    }

    public void setTwitterName(String twitterName) {
        this.twitterName = twitterName;
    }

    public boolean isVkNewsEnabled() {
        return vkNewsEnabled;
    }

    public void setVkNewsEnabled(boolean vkNewsEnabled) {
        this.vkNewsEnabled = vkNewsEnabled;
    }

    public boolean isTwitterEnabled() {
        return twitterEnabled;
    }

    public void setTwitterEnabled(boolean twitterEnabled) {
        this.twitterEnabled = twitterEnabled;
    }

    public boolean isMarkSameRead() {
        return markSameRead;
    }

    public void setMarkSameRead(boolean markSameRead) {
        this.markSameRead = markSameRead;
    }

    public boolean isUnreadOnly() {
        return unreadOnly;
    }

    public void setUnreadOnly(boolean unreadOnly) {
        this.unreadOnly = unreadOnly;
    }

    public boolean isRssEnabled() {
        return rssEnabled;
    }

    public void setRssEnabled(boolean rssEnabled) {
        this.rssEnabled = rssEnabled;
    }

    public boolean isShowPreviewButton() {
        return showPreviewButton;
    }

    public void setShowPreviewButton(boolean showPreviewButton) {
        this.showPreviewButton = showPreviewButton;
    }

    public boolean isShowTabButton() {
        return showTabButton;
    }

    public void setShowTabButton(boolean showTabButton) {
        this.showTabButton = showTabButton;
    }

    public boolean isShowReadButton() {
        return showReadButton;
    }

    public void setShowReadButton(boolean showReadButton) {
        this.showReadButton = showReadButton;
    }

}
