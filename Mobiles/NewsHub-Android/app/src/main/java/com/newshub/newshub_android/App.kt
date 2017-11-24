package com.newshub.newshub_android

import android.app.Application

import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory

class App : Application() {

    override fun onCreate() {
        super.onCreate()

//        val baseUrl = "http://192.168.0.106:3434/"
        val baseUrl = "http://13.81.71.124/"
        val retrofit = Retrofit.Builder()
                .baseUrl(baseUrl)
                .addConverterFactory(GsonConverterFactory.create())
                .build()
        api = retrofit.create<NewsHubApi>(NewsHubApi::class.java)
    }

    companion object {
        var api: NewsHubApi? = null
            private set
    }
}
