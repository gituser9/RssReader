package com.newshub.newshub_android.twitter.model

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName


data class TweetModel (
    @SerializedName("Id")
    @Expose
    val id: String,

    @SerializedName("ExpandedUrl")
    @Expose
    val expandedUrl: String,

    @SerializedName("SourceId")
    @Expose
    val sourceId: Long,

    @SerializedName("Text")
    @Expose
    val text: String
)