package com.newshub.newshub_android.twitter

import android.support.v7.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.newshub.newshub_android.R
import com.newshub.newshub_android.twitter.model.TweetModel
import com.newshub.newshub_android.twitter.model.TwitterPage
import com.newshub.newshub_android.twitter.model.TwitterSource

class TwitterViewHolder(view: View) : RecyclerView.ViewHolder(view) {
    var item: TweetModel? = null
}



class TweetRecyclerViewAdapter() : RecyclerView.Adapter<TwitterViewHolder>() {

    private var tweets: MutableList<TweetModel> = mutableListOf()
    private var sources: MutableMap<Long, TwitterSource> = mutableMapOf()


    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): TwitterViewHolder {
        val view = LayoutInflater.from(parent.context).inflate(R.layout.fragment_twitter, parent, false)
        return TwitterViewHolder(view)
    }

    override fun onBindViewHolder(holder: TwitterViewHolder, position: Int) {
        val tweet = tweets[position]
        holder.item = tweet


        /*holder.mItem = mValues.get(position)
        holder.mIdView.setText(mValues.get(position).id)
        holder.mContentView.setText(mValues.get(position).content)

        holder.mView.setOnClickListener(object : View.OnClickListener {
            public override fun onClick(v: View) {
                if (null != mListener) {
                    // Notify the active callbacks interface (the activity, if the
                    // fragment is attached to one) that an item has been selected.
                    mListener!!.onListFragmentInteraction(holder.mItem)
                }
            }
        })*/
    }

    override fun getItemCount(): Int {
        return tweets.size
    }

    fun setData(pageModel: TwitterPage) {
        tweets.clear()
        tweets.addAll(pageModel.tweets)

        pageModel.sources.forEach { source: TwitterSource ->
            sources[source.id] = source
        }

        notifyDataSetChanged()
    }

    fun addTweets(newTweets: List<TweetModel>) {
        tweets.addAll(newTweets)
        notifyDataSetChanged()
    }

    fun resetData() {
        tweets.clear()
        sources.clear()
        notifyDataSetChanged()
    }
}
