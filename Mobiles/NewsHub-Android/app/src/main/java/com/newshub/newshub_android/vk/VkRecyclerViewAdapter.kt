package com.newshub.newshub_android.vk

import android.content.Context
import android.content.Intent
import android.net.Uri
import android.support.v4.content.ContextCompat.startActivity
import android.support.v7.widget.RecyclerView
import android.text.Html
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.ImageView
import android.widget.TextView
import com.bumptech.glide.Glide
import com.newshub.newshub_android.R
import com.newshub.newshub_android.vk.model.VkGroup
import com.newshub.newshub_android.vk.model.VkNews


class VkViewHolder(view: View) : RecyclerView.ViewHolder(view) {
    var item: VkNews? = null

    val groupIcon: ImageView = view.findViewById(R.id.ivVkGroupIcon)
    val groupName: TextView = view.findViewById(R.id.tvVkGroupName)

    val newsText: TextView = view.findViewById(R.id.tvVkNewsText)
    val newsImage: ImageView = view.findViewById(R.id.ivVkImage)

    val openButton: Button = view.findViewById(R.id.btnVkOpen)
    val openLinkButton: Button = view.findViewById(R.id.btnVkOpenLink)

}


class VkRecyclerViewAdapter(var context: Context) : RecyclerView.Adapter<VkViewHolder>() {

    var groupMap: MutableMap<Long, VkGroup> = mutableMapOf()
    private var news: MutableList<VkNews> = mutableListOf()



    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): VkViewHolder {
        val view = LayoutInflater.from(parent.context).inflate(R.layout.fragment_vk_item, parent, false)
        return VkViewHolder(view)
    }

    override fun onBindViewHolder(holder: VkViewHolder, position: Int) {
        val news = news[position]
        val group = groupMap[news.groupId]
        holder.item = news

        holder.groupName.text = group?.name ?: ""
        holder.newsText.text = Html.fromHtml(news.text)

        if (!news.link.isEmpty()) {
            holder.openLinkButton.visibility = View.VISIBLE
        } else {
            holder.openLinkButton.visibility = View.GONE
        }
        if (!group?.image.isNullOrEmpty()) {
            Glide.with(context).load(group?.image).into(holder.groupIcon)
        }
        if (!news.image.isEmpty()) {
            holder.newsImage.visibility = View.VISIBLE

            if (news.image.contains(".gif?")) {
                Glide.with(context).asGif().load(news.image).into(holder.newsImage)
            } else {
                Glide.with(context).load(news.image).into(holder.newsImage)
            }

        } else {
            holder.newsImage.visibility = View.GONE
        }

        holder.openButton.setOnClickListener {
            val intent = Intent(Intent.ACTION_VIEW, Uri.parse("https://vk.com/wall-${news.groupId}_${news.postId}"))
            startActivity(context, intent, null)
        }
        holder.openLinkButton.setOnClickListener {
            if (!news.link.isEmpty()) {
                val intent = Intent(Intent.ACTION_VIEW, Uri.parse(news.link))
                startActivity(context, intent, null)
            }
        }
    }

    override fun getItemCount(): Int = news.size

    fun setData(groups: List<VkGroup>, vkNews: List<VkNews>) {
        groups.forEach { vkGroup: VkGroup ->
            groupMap.put(vkGroup.gid, vkGroup)
        }
        news.addAll(vkNews)
        notifyDataSetChanged()
    }

    fun addNews(newNews: List<VkNews>) {
        news.addAll(newNews)
        notifyDataSetChanged()
    }

}
