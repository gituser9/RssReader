package com.newshub.newshub_android


import android.support.v4.app.Fragment
import android.support.v4.app.FragmentManager
import android.support.v4.app.FragmentStatePagerAdapter
import com.newshub.newshub_android.rss.view.ContentRssFragment
import com.newshub.newshub_android.twitter.TwitterFragment
import com.newshub.newshub_android.vk.VkFragment

class PagerAdapter(fm: FragmentManager, private var mNumOfTabs: Int) : FragmentStatePagerAdapter(fm) {

    override fun getItem(position: Int): Fragment? = when (position) {
        0 -> {
            ContentRssFragment()
        }
        1 -> {
            VkFragment()
        }
        2 -> {
            TwitterFragment()
        }
        else -> null
    }

    override fun getCount(): Int = mNumOfTabs
}
