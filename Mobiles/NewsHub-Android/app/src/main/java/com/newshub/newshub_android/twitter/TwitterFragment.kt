package com.newshub.newshub_android.twitter

import android.os.Bundle
import android.support.v4.app.Fragment
import android.support.v7.widget.LinearLayoutManager
import android.support.v7.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.newshub.newshub_android.R
import com.newshub.newshub_android.general.EndlessRecyclerViewScrollListener
import com.newshub.newshub_android.twitter.model.TweetModel
import com.newshub.newshub_android.twitter.presenter.TwitterPresenter
import com.twitter.sdk.android.core.*
import com.twitter.sdk.android.core.models.Tweet
import com.twitter.sdk.android.tweetui.TweetUtils
import com.twitter.sdk.android.tweetui.TweetView
import kotlinx.android.synthetic.main.fragment_twitter_list.*


class TwitterFragment : Fragment() {

    lateinit var layoutManager: LinearLayoutManager
    lateinit var presenter: TwitterPresenter
    lateinit var adapter: TweetRecyclerViewAdapter
    var page = 1

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        Twitter.initialize(context)

        adapter = TweetRecyclerViewAdapter()
        layoutManager = LinearLayoutManager(context)
        presenter = TwitterPresenter(this)

    }

    override fun onCreateView(inflater: LayoutInflater?, container: ViewGroup?, savedInstanceState: Bundle?): View? {
        val view = inflater!!.inflate(R.layout.fragment_twitter_list, container, false)
        val recyclerView = view.findViewById<RecyclerView>(R.id.twitterRecyclerView)
        recyclerView.layoutManager = LinearLayoutManager(context)

        presenter.getTwitterPage()

        return view
    }

    override fun onViewCreated(view: View?, savedInstanceState: Bundle?) {
        val scrollListener: EndlessRecyclerViewScrollListener = object : EndlessRecyclerViewScrollListener(layoutManager) {
            override fun onLoadMore(page: Int, totalItemsCount: Int, view: RecyclerView?) {
                getTweets()
            }
        }
        twitterRecyclerView.addOnScrollListener(scrollListener)
        twitterSwiperefresh.setOnRefreshListener {
            page = 0
            adapter.resetData()
            presenter.getTwitterPage()
        }
    }

    fun showTweets(tweets: List<TweetModel>) {
        for (tweet in tweets) {
            TweetUtils.loadTweet(tweet.id.toLong(), object : Callback<Tweet>() {
                override fun success(result: Result<Tweet>?) {
                    layoutManager.addView(TweetView(context, result?.data))
                }

                override fun failure(exception: TwitterException?) {
                    TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
                }
            })
        }
    }

    private fun getTweets() {
        ++page
        presenter.getTweets(page)
        twitterSwiperefresh.isRefreshing = false
    }
}
