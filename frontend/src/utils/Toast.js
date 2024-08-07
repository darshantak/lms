import React, { useEffect } from "react";

function Toast({message,setFlag,flag}) {
    useEffect(()=>{
        if (flag){
            const timer = setTimeout(() => {
                setFlag(false)
            }, 3000);
            return () => clearTimeout(timer)
        }
    })

    if (!flag) return null;
    
  return (
    <div class="toast">
      <div class="alert alert-info">
        <span>{message}</span>
      </div>
    </div>
  );
}

export default Toast;
