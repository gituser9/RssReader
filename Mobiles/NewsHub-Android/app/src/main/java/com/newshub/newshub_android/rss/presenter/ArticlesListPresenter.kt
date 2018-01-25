package com.newshub.newshub_android.rss.presenter

import com.newshub.newshub_android.App
import com.newshub.newshub_android.general.AppSettings
import com.newshub.newshub_android.rss.model.Articles
import com.newshub.newshub_android.rss.view.ArticlesListFragment
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response


class ArticlesListPresenter(val view: ArticlesListFragment) {

    private var canLoadMore = true

    fun getArticles(feedId: Int, page: Int) {
        if (!canLoadMore && page != 1) {
            return
        }
        App.api?.getArticles(feedId, page, AppSettings.userId)?.enqueue(object : Callback<Articles> {
            override fun onResponse(call: Call<Articles>, response: Response<Articles>) {
                val articles = response.body() ?: return
                canLoadMore = articles.articles.isNotEmpty()
                println(canLoadMore)
                view.addArticles(articles.articles)
            }

            override fun onFailure(call: Call<Articles>, t: Throwable) {
                view.shortToast("Articles list load error")
            }
        })
    }

    fun markReadAll(id: Int, userId: Int) {
        App.api?.markReadAll(id, userId)?.enqueue(object : Callback<Any> {
            override fun onResponse(call: Call<Any>, response: Response<Any>) {
            }
            override fun onFailure(call: Call<Any>, t: Throwable) {
                view.shortToast("Articles list load error")
            }
        })
        view.setAllAsRead()
    }
}