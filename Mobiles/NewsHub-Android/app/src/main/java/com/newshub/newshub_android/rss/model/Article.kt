package com.newshub.newshub_android.rss.model


import java.io.Serializable
import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName

class Article : Serializable {

    @SerializedName("Id")
    @Expose
    var id: Int = 0

    @SerializedName("Title")
    @Expose
    var title: String = ""

    @SerializedName("Body")
    @Expose
    var body: String = ""

    @SerializedName("Link")
    @Expose
    var link: String = ""

    @SerializedName("IsBookmark")
    @Expose
    var isBookmark: Boolean = false


    companion object {
        private const val serialVersionUID = 3264644813100745383L
    }

}