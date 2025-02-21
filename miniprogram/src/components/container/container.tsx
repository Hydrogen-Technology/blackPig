import { View } from "@tarojs/components";
import React from "react";
import "./container.scss";

export function Container (props) {
  return (
    <View className='base-container'>   
        {props.children}
    </View>
  )
} 