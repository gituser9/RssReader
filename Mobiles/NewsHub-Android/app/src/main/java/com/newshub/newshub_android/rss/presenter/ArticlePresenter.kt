package com.newshub.newshub_android.rss.presenter

import com.newshub.newshub_android.App
import com.newshub.newshub_android.rss.model.Article
import com.newshub.newshub_android.rss.model.Articles
import com.newshub.newshub_android.rss.view.ArticleFragment
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response


class ArticlePresenter(val view: ArticleFragment) {

    fun getArticle(articleId: Int) {
        App.api?.getArticle(articleId)?.enqueue(object : Callback<Article> {
            override fun onResponse(call: Call<Article>, response: Response<Article>) {
                val article = response.body() ?: return

                view.showArticle(article)
            }

            override fun onFailure(call: Call<Article>, t: Throwable) {
                view.shortToast("Article load error")
            }
        })
    }

}