import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";

function Toast({message,setFlag,flag}) {
  const navigate = useNavigate();
    useEffect(()=>{
        if (flag){
            const timer = setTimeout(() => {
                setFlag(false)
                navigate("/dashboard");
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
