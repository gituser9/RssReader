package com.newshub.newshub_android.rss.model

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName

import java.io.Serializable

class Feed : Serializable {

    // TODO: add long int

    @SerializedName("Id")
    @Expose
    var id: Int = 0
    @SerializedName("Name")
    @Expose
    var name: String? = null
    @SerializedName("Url")
    @Expose
    var url: String? = null
    @SerializedName("UserId")
    @Expose
    var userId: Int = 0
    @SerializedName("Articles")
    @Expose
    var articles: Any? = null

}
