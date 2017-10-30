package com.newshub.newshub_android.rss.view

import android.support.v4.app.Fragment
import android.widget.Toast


abstract class BaseRssFragment : Fragment() {

    fun shortToast(message: String) {
        Toast.makeText(activity, message, Toast.LENGTH_SHORT).show()
    }

}