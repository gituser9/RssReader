package com.newshub.newshub_android.twitter.model

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName


data class TwitterSource (
        @SerializedName("Id")
        @Expose
        val id: Long,

        @SerializedName("Image")
        @Expose
        val image: String,

        @SerializedName("Name")
        @Expose
        val name: String,

        @SerializedName("ScreenName")
        @Expose
        val screenName: String,

        @SerializedName("Url")
        @Expose
        val url: String
)