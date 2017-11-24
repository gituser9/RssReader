package com.newshub.newshub_android.rss.view.adapter

import android.graphics.Typeface
import android.support.v7.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView

import com.newshub.newshub_android.R
import com.newshub.newshub_android.rss.view.FeedListFragment.OnListFragmentInteractionListener
import com.newshub.newshub_android.rss.model.Feed
import com.newshub.newshub_android.rss.model.FeedModel

import java.util.ArrayList

class FeedListViewHolder(var view: View) : RecyclerView.ViewHolder(view) {
    lateinit var mItem: Feed
    val mContentView: TextView = view.findViewById(R.id.content)
}


class FeedListRecyclerViewAdapter() : RecyclerView.Adapter<FeedListViewHolder>() {

    var feeds: List<FeedModel> = emptyList()
    lateinit var listener: OnListFragmentInteractionListener


    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): FeedListViewHolder {
        val view = LayoutInflater.from(parent.context).inflate(R.layout.fragment_feedlist, parent, false)

        return FeedListViewHolder(view)
    }

    override fun onBindViewHolder(holder: FeedListViewHolder, position: Int) {
        val feedModel = feeds[position]
        val context = holder.view.context
        holder.mItem = feedModel.feed!!

        if (feedModel.isExistUnread) {
            holder.mContentView.text = context.getString(R.string.feed_row, feedModel.feed?.name, feedModel.articlesCount.toString())
            holder.mContentView.typeface = Typeface.DEFAULT_BOLD
        } else {
            holder.mContentView.text = feedModel.feed?.name
            holder.mContentView.typeface = Typeface.DEFAULT
        }

        holder.view.setOnClickListener { _: View? ->
            listener.onListFragmentInteraction(feedModel)
        }
    }

    override fun getItemCount(): Int {
        return feeds.size
    }

}
