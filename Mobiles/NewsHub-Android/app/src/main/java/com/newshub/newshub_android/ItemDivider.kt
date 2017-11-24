package com.newshub.newshub_android

import android.content.Context
import android.graphics.Canvas
import android.graphics.drawable.Drawable
import android.support.v7.widget.RecyclerView
import android.view.View


class ItemDivider// Конструктор загружает встроенный разделитель элементов списка
(context: Context) : RecyclerView.ItemDecoration() {
    private val divider: Drawable?

    init {
        val attrs = intArrayOf(android.R.attr.listDivider)
        divider = context.obtainStyledAttributes(attrs).getDrawable(0)
    }

    // Рисование разделителей элементов списка в RecyclerView
    override fun onDrawOver(c: Canvas, parent: RecyclerView, state: RecyclerView.State?) {
        super.onDrawOver(c, parent, state)

        // Вычисление координат x для всех разделителей
        val left = parent.paddingLeft
        val right = parent.width - parent.paddingRight

        // Для каждого элемента, кроме последнего, нарисовать линию
        for (i in 0 until parent.childCount) {
            val item = parent.getChildAt(i)

            // Вычисление координат y текущего разделителя
            val top = item.bottom + (item.layoutParams as RecyclerView.LayoutParams).bottomMargin
            val bottom = top + divider!!.intrinsicHeight

            // Рисование разделителя с вычисленными границами
            divider.setBounds(left, top, right, bottom)
            divider.draw(c)
        }
    }


}
