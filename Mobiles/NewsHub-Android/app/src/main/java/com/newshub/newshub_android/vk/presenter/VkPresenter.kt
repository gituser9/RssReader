package com.newshub.newshub_android.vk.presenter

import com.newshub.newshub_android.App
import com.newshub.newshub_android.general.AppSettings
import com.newshub.newshub_android.vk.VkFragment
import com.newshub.newshub_android.vk.model.VkNews
import com.newshub.newshub_android.vk.model.VkPage
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response


class VkPresenter(val view: VkFragment) {

    fun getPage() {
        App.api?.getVkPage(AppSettings.userId)?.enqueue(object : Callback<VkPage> {
            override fun onResponse(call: Call<VkPage>, response: Response<VkPage>) {
                val page = response.body() ?: return

                view.adapter.setData(page.groups, page.news)
            }

            override fun onFailure(call: Call<VkPage>, t: Throwable) {

            }
        })
    }

    fun getNews(page: Int) {
        App.api?.getVkNews(AppSettings.userId, page)?.enqueue(object : Callback<List<VkNews>> {
            override fun onResponse(call: Call<List<VkNews>>, response: Response<List<VkNews>>) {
                val news = response.body() ?: return

                view.adapter.addNews(news)
            }

            override fun onFailure(call: Call<List<VkNews>>, t: Throwable) {

            }
        })
    }

}