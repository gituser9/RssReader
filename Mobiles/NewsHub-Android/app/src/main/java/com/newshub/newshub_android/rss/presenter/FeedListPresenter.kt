package com.newshub.newshub_android.rss.presenter

import android.widget.Toast
import com.newshub.newshub_android.App
import com.newshub.newshub_android.general.AppSettings
import com.newshub.newshub_android.rss.model.Feed
import com.newshub.newshub_android.rss.model.FeedModel
import com.newshub.newshub_android.rss.view.FeedListFragment
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response
import java.util.ArrayList


class FeedListPresenter(val view: FeedListFragment) {

    fun getAll() {
        App.api?.getAllFeeds(AppSettings.userId)?.enqueue(object : Callback<List<FeedModel>> {
            override fun onResponse(call: Call<List<FeedModel>>, response: Response<List<FeedModel>>) {
                val feedModels = response.body() ?: return

                view.showAll(feedModels)
            }

            override fun onFailure(call: Call<List<FeedModel>>, t: Throwable) {
                view.shortToast("Feed list load error")
            }
        })
    }

}