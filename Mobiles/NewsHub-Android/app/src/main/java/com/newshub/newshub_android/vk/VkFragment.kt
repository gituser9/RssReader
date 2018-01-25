package com.newshub.newshub_android.vk

import android.os.Bundle
import android.support.v4.app.Fragment
import android.support.v7.widget.LinearLayoutManager
import android.support.v7.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.newshub.newshub_android.R
import com.newshub.newshub_android.general.EndlessRecyclerViewScrollListener
import com.newshub.newshub_android.vk.presenter.VkPresenter
import kotlinx.android.synthetic.main.fragment_vk_item_list.*


class VkFragment : Fragment() {

    lateinit var adapter: VkRecyclerViewAdapter
    lateinit var presenter: VkPresenter
    lateinit var layoutManager: LinearLayoutManager
    var page = 1


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        adapter = VkRecyclerViewAdapter(context)
        presenter = VkPresenter(this)
        layoutManager = LinearLayoutManager(context)
    }

    override fun onCreateView(inflater: LayoutInflater?, container: ViewGroup?, savedInstanceState: Bundle?): View? {
        val view = inflater!!.inflate(R.layout.fragment_vk_item_list, container, false)

        val recyclerView = view.findViewById<RecyclerView>(R.id.vkRecyclerView)
        recyclerView.adapter = adapter
        recyclerView.layoutManager = layoutManager

        presenter.getPage()

        return view
    }

    override fun onViewCreated(view: View?, savedInstanceState: Bundle?) {
        val scrollListener: EndlessRecyclerViewScrollListener = object : EndlessRecyclerViewScrollListener(layoutManager) {
            override fun onLoadMore(page: Int, totalItemsCount: Int, view: RecyclerView?) {
                getNews()
            }
        }
        vkRecyclerView.addOnScrollListener(scrollListener)
        vkSwiperefresh.setOnRefreshListener {
            page = 0
            adapter.resetNews()
            getNews()
        }
    }

    private fun getNews() {
        ++page
        presenter.getNews(page)
        vkSwiperefresh.isRefreshing = false
    }
}
