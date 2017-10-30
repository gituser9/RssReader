package com.newshub.newshub_android.rss.view

import android.os.Bundle
import android.support.v4.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.newshub.newshub_android.R


class ContentRssFragment : Fragment() {

    override fun onCreateView(inflater: LayoutInflater?, container: ViewGroup?, savedInstanceState: Bundle?): View? {
        val view = inflater!!.inflate(R.layout.fragment_content_rss, container, false)
        val fragment = FeedListFragment()
        val transaction = activity.supportFragmentManager.beginTransaction()
        transaction.replace(R.id.content_rss, fragment)
        transaction.addToBackStack(null)
        transaction.commit()

        return view
    }

}
