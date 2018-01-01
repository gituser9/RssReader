package com.newshub.newshub_android.vk.model

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName


data class VkGroup(
        @SerializedName("Id")
        @Expose
        val id: Long,

        @SerializedName("Gid")
        @Expose
        val gid: Long,

        @SerializedName("Image")
        @Expose
        val image: String,

        @SerializedName("LinkedName")
        @Expose
        val linkedName: String,

        @SerializedName("Name")
        @Expose
        val name: String
)