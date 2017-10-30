package com.newshub.newshub_android.rss.view.adapter

import android.graphics.Typeface
import android.support.v7.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView

import com.newshub.newshub_android.R
import com.newshub.newshub_android.rss.view.ArticlesListFragment.OnListFragmentInteractionListener
import com.newshub.newshub_android.rss.model.ArticleTitle

import java.util.ArrayList


class ArticlesListViewHolder(var mView: View) : RecyclerView.ViewHolder(mView) {
    val mContentView: TextView = mView.findViewById(R.id.content)

    var mItem: ArticleTitle? = null

}

class ArticlesListRecyclerViewAdapter : RecyclerView.Adapter<ArticlesListViewHolder>() {

    val articles = mutableListOf<ArticleTitle>()
    var listener: OnListFragmentInteractionListener? = null


    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ArticlesListViewHolder {
        val view = LayoutInflater.from(parent.context).inflate(R.layout.fragment_articleslist, parent, false)
        return ArticlesListViewHolder(view)
    }

    override fun onBindViewHolder(holder: ArticlesListViewHolder, position: Int) {
        val articleTitle = articles[position]
        holder.mItem = articleTitle
        holder.mContentView.text = articleTitle.title

        if (!articleTitle.isRead) {
            holder.mContentView.typeface = Typeface.DEFAULT_BOLD
        }

        holder.mView.setOnClickListener { _: View? ->
            listener?.onListFragmentInteraction(articleTitle)
        }
    }

    override fun getItemCount(): Int {
        return articles.size
    }

    fun setArticles(articles: List<ArticleTitle>) {
        this.articles.addAll(articles)
    }
}
