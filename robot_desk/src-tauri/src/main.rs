// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use serde::{Deserialize, Serialize};
use tauri::Window;

#[derive(Debug, Serialize, Deserialize)]
struct Meta {
    code: i32,
    message: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct ApiResponse {
    meta: Meta,
    total: Option<i64>,
    data: Option<serde_json::Value>,
}

#[derive(Clone, serde::Serialize)]
struct Payload {
    message: String,
}


// --------------------------login-----------------------
#[derive(Debug, Deserialize, Serialize)]
struct LoginRequest {
    username: String,
    password: String,
}

#[tauri::command]
async fn login(window: Window, request: LoginRequest) -> Option<ApiResponse> {
    match reqwest::Client::new()
        .post("http://121.36.209.8:8080/user/login")
        .header("Content-Type", "application/json")
        .json(&request)
        .send()
        .await
    {
        Ok(response) => {
            // 处理成功的响应
            match response.json().await {
                Ok(json_res) => Some(json_res),
                Err(_) => {
                    window.emit("network-error", Payload { message: "数据解析异常".into() }).unwrap();
                    None
                }
            }
        }
        Err(_) => {
            // 处理错误，发送事件并返回 None
            window
                .emit("network-error", Payload {
                    message: "网络出现错误，请检查本地网络".into(),
                })
                .unwrap();
            None
        }
    }
}

#[derive(Debug, Serialize, Deserialize)]
struct PaginationParam {
    page: usize,
    page_size: usize
}

#[tauri::command]
async fn fetch_robots(window: Window, paginator: PaginationParam) -> Option<ApiResponse> {
    match reqwest::Client::new().post("http://121.36.209.8:8080/robot/get_by_page")
        .json(&paginator)
        .send()
        .await {
        Ok(res) => {
            match res.json().await {
                Ok(json_res) => Some(json_res),
                Err(_) => {
                    window.emit("network-error", Payload { message: "数据解析异常".into() }).unwrap();
                    None
                }
            }
        }
        Err(_) => {
            window.emit("network-error", Payload { message: "网络出现错误，请检查本地网络".into() }).unwrap();
            None
        }
    }
}

// 用户注册
#[tauri::command]
async fn create_user(window: Window, register: LoginRequest) -> Option<ApiResponse> {
    match reqwest::Client::new().post("http://121.36.209.8:8080/user/add")
        .header("Content-Type", "application/json")
        .json(&register)
        .send()
        .await {
        Ok(res) => {
            match res.json().await {
                Ok(json_res) => Some(json_res),
                Err(_) => {
                    window.emit("network-error", Payload { message: "数据解析异常".into() }).unwrap();
                    None
                }
            }
        }
        Err(_) => {
            window.emit("network-error", Payload { message: "网络出现错误，请检查本地网络".into() }).unwrap();
            None
        }
    }
}

#[tokio::main]
async fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![fetch_robots, login,create_user])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
