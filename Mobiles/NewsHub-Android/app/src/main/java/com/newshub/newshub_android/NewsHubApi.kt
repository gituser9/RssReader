package com.newshub.newshub_android

import com.newshub.newshub_android.general.LoginData
import com.newshub.newshub_android.general.User
import com.newshub.newshub_android.rss.model.Article
import com.newshub.newshub_android.rss.model.Articles
import com.newshub.newshub_android.rss.model.FeedModel
import com.newshub.newshub_android.settings.model.Settings
import com.newshub.newshub_android.vk.model.VkNews
import com.newshub.newshub_android.vk.model.VkPage
import retrofit2.Call
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.POST
import retrofit2.http.Query

interface NewsHubApi {

    // General
    @POST("auth")
    fun auth(@Body data: LoginData): Call<User>

    // Rss
    @GET("get-all")
    fun getAllFeeds(@Query("id") id: Int): Call<List<FeedModel>>

    @GET("get-settings")
    fun getSettings(@Query("id") id: Int): Call<Settings>

    @GET("get-articles")
    fun getArticles(@Query("id") id: Int, @Query("page") page: Int, @Query("userId") userId: Int): Call<Articles>

    @GET("get-article")
    fun getArticle(@Query("id") id: Int): Call<Article>

    @GET("mark-read-all")
    fun markReadAll(@Query("id") id: Int, @Query("userId") userId: Int): Call<Any>

    // Vk
    @GET("get-vk-news")
    fun getVkNews(@Query("id") id: Int, @Query("page") page: Int): Call<List<VkNews>>

    @GET("get-vk-page")
    fun getVkPage(@Query("id") id: Int): Call<VkPage>
}
