package com.newshub.newshub_android.vk.model

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName


class VkPage(
        @SerializedName("Groups")
        @Expose
        val groups: List<VkGroup>,

        @SerializedName("News")
        @Expose
        val news: List<VkNews>
)