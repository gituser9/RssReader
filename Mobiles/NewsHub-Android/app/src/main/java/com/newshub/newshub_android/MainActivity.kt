package com.newshub.newshub_android

import android.content.Context
import android.content.Intent
import android.support.design.widget.TabLayout
import android.support.v4.view.ViewPager
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.view.View
import android.widget.ProgressBar
import android.widget.Toast

import com.newshub.newshub_android.general.AppSettings
import com.newshub.newshub_android.settings.model.Settings

import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class MainActivity : AppCompatActivity() {
    lateinit var mainProgressBar: ProgressBar
    lateinit var tabLayout: TabLayout

    private var userId: Int = 0


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        getPreferences()

        if (userId == 0) {
            // login
            showLoginActivity()
        }

        AppSettings.userId = userId
        start()
    }

    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        if (data == null) {
            showLoginActivity()
        }

        getPreferences()
        start()
    }

    private fun start() {
        AppSettings.userId = userId
        mainProgressBar = findViewById(R.id.mainProgressBar)
        tabLayout = findViewById(R.id.tabLayout)

        App.api?.getSettings(userId)?.enqueue(object : Callback<Settings> {
            override fun onResponse(call: Call<Settings>, response: Response<Settings>) {
                settings = response.body()
                //                posts.addAll(response.body());
                //                recyclerView.getAdapter().notifyDataSetChanged();
                mainProgressBar.visibility = View.GONE
                createTabs(settings)
                setupTabLayout()
            }

            override fun onFailure(call: Call<Settings>, t: Throwable) {
                //                System.out.println(t.getMessage());
                mainProgressBar.visibility = View.GONE
                Toast.makeText(this@MainActivity, "An error occurred during networking", Toast.LENGTH_SHORT).show()
            }
        })
    }

    private fun showLoginActivity() {
        val intent = Intent(this, LoginActivity::class.java)
        startActivity(intent)
        return
    }

    private fun saveSettings(settings: Settings) {

    }

    private fun getPreferences() {
        val preferences = getSharedPreferences("userInfo", Context.MODE_PRIVATE)
        userId = preferences.getInt("userId", 0)
    }

    private fun createTabs(settings: Settings?) {
        var sourceNum = 0

        if (settings!!.isRssEnabled) {
            tabLayout.addTab(tabLayout.newTab().setText("Rss"))
            ++sourceNum
        }
        /*if (settings.isVkNewsEnabled()) {
            tabLayout.addTab(tabLayout.newTab().setText("Vk"));
            ++sourceNum;
        }
        if (settings.isTwitterEnabled()) {
            tabLayout.addTab(tabLayout.newTab().setText("Twitter"));
            ++sourceNum;
        }*/
        if (sourceNum < 2) {
            tabLayout.visibility = View.GONE
        }

        tabLayout.tabGravity = TabLayout.GRAVITY_FILL
    }

    private fun setupTabLayout() {
        val viewPager = findViewById<ViewPager>(R.id.viewPager)
        val adapter = PagerAdapter(supportFragmentManager, tabLayout.tabCount)
        viewPager.adapter = adapter
        viewPager.addOnPageChangeListener(TabLayout.TabLayoutOnPageChangeListener(tabLayout))

        tabLayout.setOnTabSelectedListener(object : TabLayout.OnTabSelectedListener {
            override fun onTabSelected(tab: TabLayout.Tab) {
                viewPager.currentItem = tab.position
            }

            override fun onTabUnselected(tab: TabLayout.Tab) {

            }

            override fun onTabReselected(tab: TabLayout.Tab) {

            }
        })
    }

    companion object {

        var settings: Settings? = null
    }
}
