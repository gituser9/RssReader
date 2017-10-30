package com.newshub.newshub_android


import android.support.v4.app.Fragment
import android.support.v4.app.FragmentManager
import android.support.v4.app.FragmentStatePagerAdapter
import com.newshub.newshub_android.rss.view.ContentRssFragment

class PagerAdapter(fm: FragmentManager, internal var mNumOfTabs: Int) : FragmentStatePagerAdapter(fm) {

    override fun getItem(position: Int): Fragment? {

        when (position) {
            0 -> {
                return ContentRssFragment()
            }
        /*case 1:
                TabFragment2 tab2 = new TabFragment2();
                return tab2;
            case 2:
                TabFragment3 tab3 = new TabFragment3();
                return tab3;*/
            else -> return null
        }
    }

    override fun getCount(): Int {
        return mNumOfTabs
    }
}