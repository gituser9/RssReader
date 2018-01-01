package com.newshub.newshub_android.presenter

import com.newshub.newshub_android.NewsHubApi
import com.newshub.newshub_android.rss.model.FeedModel
import org.junit.Test
import org.junit.runner.RunWith
import org.mockito.Mockito.`when`
import org.mockito.Mockito.mock
import org.robolectric.RobolectricTestRunner
import retrofit2.Call


@RunWith(RobolectricTestRunner::class)
class FeedListPresenterTest  {

    @Test
    fun shouldGetFeedList() {
        val mockApi = mock(NewsHubApi::class.java)
        `when`(mockApi.getAllFeeds(0)).thenReturn(Call<List<FeedModel>>)


    }

}