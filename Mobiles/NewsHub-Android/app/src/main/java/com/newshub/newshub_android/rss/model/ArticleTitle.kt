package com.newshub.newshub_android.rss.model

import java.io.Serializable;
import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName


class ArticleTitle : Serializable {

    @SerializedName("Id")
    @Expose
    var id: Int = 0

    @SerializedName("Title")
    @Expose
    var title: String = ""

    @SerializedName("IsRead")
    @Expose
    var isRead: Boolean = false

    @SerializedName("IsBookmark")
    @Expose
    var isBookmark: Boolean = false

    companion object {
        private val serialVersionUID = 3102038750294094453L
    }

}

class Articles : Serializable {

    @SerializedName("Articles")
    @Expose
    var articles: List<ArticleTitle> = emptyList()

    @SerializedName("Count")
    @Expose
    var count: Int = 0

    companion object {
        private val serialVersionUID = -6267315743399763135L
    }

}