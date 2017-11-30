package com.newshub.newshub_android.rss.view

import android.os.Bundle
import android.support.v7.widget.LinearLayoutManager
import android.support.v7.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.newshub.newshub_android.ItemDivider
import com.newshub.newshub_android.R
import com.newshub.newshub_android.rss.model.FeedModel
import com.newshub.newshub_android.rss.presenter.FeedListPresenter
import com.newshub.newshub_android.rss.view.adapter.FeedListRecyclerViewAdapter


class FeedListFragment : BaseRssFragment() {

    private var mListener: OnListFragmentInteractionListener? = null
    private lateinit var adapter: FeedListRecyclerViewAdapter
    private lateinit var presenter: FeedListPresenter


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        presenter = FeedListPresenter(this)
        adapter = FeedListRecyclerViewAdapter()

        // get feeds
        presenter.getAll()
    }

    override fun onCreateView(inflater: LayoutInflater?, container: ViewGroup?, savedInstanceState: Bundle?): View? {
        val view = inflater!!.inflate(R.layout.fragment_feedlist_list, container, false)
        val recyclerView = view.findViewById<RecyclerView>(R.id.feedList)

        // Set the adapter
        recyclerView.layoutManager = LinearLayoutManager(activity.baseContext)
        recyclerView.adapter = adapter
        recyclerView.addItemDecoration(ItemDivider(view.context))
        recyclerView.setHasFixedSize(true)
        recyclerView.isClickable = true

        return view
    }


    fun showAll(feeds: List<FeedModel>) {
        adapter.listener = object : OnListFragmentInteractionListener {
            override fun onListFragmentInteraction(item: FeedModel) {
                val bundle = Bundle()
//                bundle.putSerializable(AppSettings.feedKey, item)

                val fragment = ArticlesListFragment()
                fragment.arguments = bundle
                fragment.feedModel = item

                val transaction = activity.supportFragmentManager.beginTransaction()
                transaction.replace(R.id.content_rss, fragment)
                transaction.addToBackStack("Feeds")
                transaction.commit()
            }
        }
        adapter.feeds = feeds
        adapter.notifyDataSetChanged()
    }

    override fun onDetach() {
        super.onDetach()
        mListener = null
    }


    interface OnListFragmentInteractionListener {
        fun onListFragmentInteraction(item: FeedModel)
    }

}


