package com.newshub.newshub_android.vk.model

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName


data class VkNews(
        @SerializedName("Id")
        @Expose
        val id: Long,

        @SerializedName("GroupId")
        @Expose
        val groupId: Long,

        @SerializedName("Image")
        @Expose
        val image: String,

        @SerializedName("PostId")
        @Expose
        val postId: Long,

        @SerializedName("Text")
        @Expose
        val text: String,

        @SerializedName("Link")
        @Expose
        val link: String
)