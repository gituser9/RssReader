package com.newshub.newshub_android.rss.view

import android.os.Bundle
import android.support.v7.widget.LinearLayoutManager
import android.support.v7.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup

import com.newshub.newshub_android.ItemDivider
import com.newshub.newshub_android.R
import com.newshub.newshub_android.general.AppSettings
import com.newshub.newshub_android.general.EndlessRecyclerViewScrollListener
import com.newshub.newshub_android.rss.model.ArticleTitle
import com.newshub.newshub_android.rss.model.FeedModel
import com.newshub.newshub_android.rss.presenter.ArticlesListPresenter
import com.newshub.newshub_android.rss.view.adapter.ArticlesListRecyclerViewAdapter
import kotlinx.android.synthetic.main.fragment_articleslist_list.*


class ArticlesListFragment : BaseRssFragment() {

    lateinit var feedModel: FeedModel
    private lateinit var presenter: ArticlesListPresenter
    private lateinit var layoutManager: LinearLayoutManager
    private lateinit var recyclerView: RecyclerView
    private var mListener: OnListFragmentInteractionListener? = null
    private var page = 0
    private val adapter = ArticlesListRecyclerViewAdapter()


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        activity.title = feedModel.feed.name
        presenter = ArticlesListPresenter(this)
        adapter.listener = object : OnListFragmentInteractionListener {
            override fun onListFragmentInteraction(item: ArticleTitle) {
                item.isRead = true

                val bundle = Bundle()
                bundle.putInt(AppSettings.articleKey, item.id)

                val fragment = ArticleFragment()
                fragment.arguments = bundle
                fragment.articleId = item.id

                if (feedModel.articlesCount > 0) {
                    --feedModel.articlesCount
                    feedModel.isExistUnread = feedModel.articlesCount > 0
                }

                val transaction = activity.supportFragmentManager.beginTransaction()
                transaction.replace(R.id.content_rss, fragment)
                transaction.addToBackStack("ArticleList")
                transaction.commit()
            }
        }

        getArticles()
    }

    override fun onCreateView(inflater: LayoutInflater?, container: ViewGroup?, savedInstanceState: Bundle?): View? {
        val view = inflater!!.inflate(R.layout.fragment_articleslist_list, container, false)
        recyclerView = view.findViewById(R.id.articles_list_recycler_view)

        // Set the adapter
        val context = view.getContext()

        layoutManager = LinearLayoutManager(context)
        recyclerView.layoutManager = layoutManager
        recyclerView.adapter = adapter
        recyclerView.addItemDecoration(ItemDivider(context))

        return view
    }

    override fun onViewCreated(view: View?, savedInstanceState: Bundle?) {
        val scrollListener: EndlessRecyclerViewScrollListener = object : EndlessRecyclerViewScrollListener(layoutManager) {
            override fun onLoadMore(page: Int, totalItemsCount: Int, view: RecyclerView?) {
                getArticles()
            }
        }
        recyclerView.addOnScrollListener(scrollListener)

        btnMarkReadAll.setOnClickListener {
            presenter.markReadAll(feedModel.feed.id, AppSettings.userId)
        }
    }

    override fun onDetach() {
        super.onDetach()
        mListener = null
    }

    interface OnListFragmentInteractionListener {
        fun onListFragmentInteraction(item: ArticleTitle)
    }

    private fun getArticles() {
        ++page
        presenter.getArticles(feedModel.feed.id, page)
    }

    fun addArticles(articles: List<ArticleTitle>) {
        adapter.setArticles(articles)
        adapter.notifyDataSetChanged()
    }

    fun setAllAsRead() {
        adapter.articles.forEach { articleTitle: ArticleTitle ->
            articleTitle.isRead = true
        }
        adapter.notifyItemRangeChanged(0, adapter.articles.size)
        feedModel.isExistUnread = false
    }
}
