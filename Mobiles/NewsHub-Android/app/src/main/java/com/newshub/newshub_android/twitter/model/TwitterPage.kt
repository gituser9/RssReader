package com.newshub.newshub_android.twitter.model

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName


data class TwitterPage (
        @SerializedName("News")
        @Expose
        val tweets: List<TweetModel>,

        @SerializedName("Sources")
        @Expose
        val sources: List<TwitterSource>
)