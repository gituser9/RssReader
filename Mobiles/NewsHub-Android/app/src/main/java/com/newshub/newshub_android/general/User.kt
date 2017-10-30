package com.newshub.newshub_android.general


import java.io.Serializable
import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName

class User : Serializable {

    @SerializedName("Id")
    @Expose
    var id: Int = 0

    @SerializedName("Name")
    @Expose
    var name: String = ""

    @SerializedName("Password")
    @Expose
    var password: String = ""

    @SerializedName("VkLogin")
    @Expose
    var vkLogin: String = ""

    @SerializedName("VkPassword")
    @Expose
    var vkPassword: String = ""

    @SerializedName("TwitterScreenName")
    @Expose
    var twitterScreenName: String = ""

    @SerializedName("VkNewsEnabled")
    @Expose
    var vkNewsEnabled: Boolean = false

    @SerializedName("Settings")
    @Expose
    var settings: Settings = Settings()


    companion object {
        private const val serialVersionUID = 337225686894046188L
    }

}


class Settings : Serializable {

    @SerializedName("Id")
    @Expose
    var id: Int = 0

    @SerializedName("UserId")
    @Expose
    var userId: Int = 0

    @SerializedName("UnreadOnly")
    @Expose
    var unreadOnly: Boolean = false

    @SerializedName("MarkSameRead")
    @Expose
    var markSameRead: Boolean = false

    @SerializedName("RssEnabled")
    @Expose
    var rssEnabled: Boolean = false

    @SerializedName("VkNewsEnabled")
    @Expose
    var vkNewsEnabled: Boolean = false

    @SerializedName("TwitterEnabled")
    @Expose
    var twitterEnabled: Boolean = false

    @SerializedName("ShowPreviewButton")
    @Expose
    var showPreviewButton: Boolean = false

    @SerializedName("ShowTabButton")
    @Expose
    var showTabButton: Boolean = false

    @SerializedName("ShowReadButton")
    @Expose
    var showReadButton: Boolean = false


    companion object {
        private const val serialVersionUID = 1660006493847197451L
    }

}