package com.newshub.newshub_android.rss.model


import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName
import com.newshub.newshub_android.rss.model.Feed

class FeedModel {

    @SerializedName("Feed")
    @Expose
    var feed: Feed = Feed()

    @SerializedName("ArticlesCount")
    @Expose
    var articlesCount: Int = 0

    @SerializedName("ExistUnread")
    @Expose
    var isExistUnread: Boolean = false

}


