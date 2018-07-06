package com.newshub.newshub_android.twitter.presenter

import com.newshub.newshub_android.App
import com.newshub.newshub_android.general.AppSettings
import com.newshub.newshub_android.twitter.TwitterFragment
import com.newshub.newshub_android.twitter.model.TweetModel
import com.newshub.newshub_android.twitter.model.TwitterPage
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response


class TwitterPresenter(val view: TwitterFragment) {
    fun getTwitterPage() {
        App.api?.getTwitterPage(AppSettings.userId)?.enqueue(object : Callback<TwitterPage> {
            override fun onResponse(call: Call<TwitterPage>?, response: Response<TwitterPage>?) {
                val page = response?.body() ?: return
                view.adapter.setData(page)
            }

            override fun onFailure(call: Call<TwitterPage>?, t: Throwable?) {

            }
        })
    }

    fun getTweets(page: Int) {
        App.api?.getTweets(AppSettings.userId, page)?.enqueue(object : Callback<List<TweetModel>> {
            override fun onResponse(call: Call<List<TweetModel>>?, response: Response<List<TweetModel>>?) {
                val tweets = response?.body() ?: return

                view.adapter.addTweets(tweets)
            }

            override fun onFailure(call: Call<List<TweetModel>>?, t: Throwable?) {

            }
        })
    }
}