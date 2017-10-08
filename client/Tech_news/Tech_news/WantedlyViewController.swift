//
//  WantedlyViewController.swift
//  Tech_news
//
//  Created by 栗栖遼馬 on 2017/10/05.
//  Copyright © 2017年 Ryoma. All rights reserved.
//

import UIKit

import Foundation

class WantedlyViewController: UIViewController,UITableViewDelegate,UITableViewDataSource,UIWebViewDelegate {
    
    var tableView:UITableView = UITableView()
    var refleshControl:UIRefreshControl!
    var webView:UIWebView = UIWebView()
    var goButton:UIButton!
    var backButton:UIButton!
    var cancelButton:UIButton!
    var dotsView:DotsLoader! = DotsLoader()
    let decoder: JSONDecoder = JSONDecoder()
    var articles: [Article]?
    var elements = NSMutableDictionary()
    var element = String()
    var text = NSMutableString()
    
    
    //json struct
    struct Article : Codable{
        let id:Int
        let title:String
        let posted_at:String
        let url:String
        let text:String
        let company:String
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        //引っ張って更新
        refleshControl = UIRefreshControl()
        refleshControl.tintColor = UIColor.white
        refleshControl.addTarget(self, action: #selector(reflesh), for: UIControlEvents.valueChanged)
        
        //tableview
        tableView.delegate = self
        tableView.dataSource = self
        tableView.frame = CGRect(x: 0, y: 0, width: view.frame.size.width, height: view.frame.size.height - 54.0)
        tableView.backgroundColor = UIColor.clear
        tableView.addSubview(refleshControl)
        self.view.addSubview(tableView)
        
        //webview
        webView.frame = tableView.frame
        webView.delegate = self
        webView.scalesPageToFit = true
        webView.contentMode = .scaleAspectFit
        self.view.addSubview(webView)
        webView.isHidden = true
        
        //1つ進むボタン
        goButton = UIButton()
        goButton.frame = CGRect(x: self.view.frame.size.width - 50, y: self.view.frame.size.height - 128, width: 50, height: 50)
        goButton.setImage(UIImage(named:"go.png"), for: .normal)
        goButton.addTarget(self, action: #selector(nextPage), for: .touchUpInside)
        self.view.addSubview(goButton)
        
        //戻るボタン
        backButton = UIButton()
        backButton.frame = CGRect(x:  10, y: self.view.frame.size.height - 128, width: 50, height: 50)
        backButton.setImage(UIImage(named:"back.png"), for: .normal)
        backButton.addTarget(self, action: #selector(backPage), for: .touchUpInside)
        self.view.addSubview(backButton)
        
        //キャンセルボタン
        cancelButton = UIButton()
        cancelButton.frame = CGRect(x:  10, y: 80, width: 50, height: 50)
        cancelButton.setImage(UIImage(named:"cancel.png"), for: .normal)
        cancelButton.addTarget(self, action: #selector(cancelPage), for: .touchUpInside)
        self.view.addSubview(cancelButton)
        
        goButton.isHidden = true
        backButton.isHidden = true
        cancelButton.isHidden = true
        
        //ドッツビュー
        dotsView.frame = CGRect(x: 0, y: self.view.frame.size.height/3, width: self.view.frame.size.width, height: 100)
        dotsView.dotsCount = 5
        dotsView.dotsRadius = 10
        self.view.addSubview(dotsView)
        
        dotsView.isHidden = true
        
        //jsonパース
        let url:String =   address + "/api/articles/?company=wantedly"
        let urltoSend:URL = URL(string:url)!
        articles = []
        do {
            let data = try Data(contentsOf: urltoSend, options: [])
            articles = try JSONDecoder().decode([Article].self, from: data)
        }catch{
            print(error)
        }
        tableView.reloadData()
        
        
        
    }
    
    
    @objc func backPage(){
        webView.goBack()
    }
    
    @objc func cancelPage(){
        webView.isHidden = true
        goButton.isHidden  = true
        cancelButton.isHidden = true
        backButton.isHidden = true
    }
    
    @objc func nextPage(){
        webView.goForward()
    }
    
    @objc func reflesh(){
        perform(#selector(delay), with: nil, afterDelay: 2.0)
        
    }
    
    @objc func delay(){
        let url:String =  address + "/api/articles/?company=wantedly"
        let urltoSend:URL = URL(string:url)!
        articles = []
        do {
            let data = try Data(contentsOf: urltoSend, options: [])
            articles = try JSONDecoder().decode([Article].self, from: data)
        }catch{
            print(error)
        }
        tableView.reloadData()
        refleshControl.endRefreshing()
    }
    
    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return 100
    }
    
    func numberOfSections(in tableView: UITableView) -> Int {
        return 1
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return articles!.count
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = UITableViewCell(style: .subtitle, reuseIdentifier: "Cell")
        
        cell.selectionStyle = .none
        //cell.backgroundColor= UIColor.clear
        //改行消さないと表示できない
        cell.textLabel?.text = String(articles![indexPath.row].title.filter { !"\n\r".contains($0) })
        cell.textLabel?.font = UIFont.boldSystemFont(ofSize: 15.0)
        cell.textLabel?.textColor = UIColor.black
        
        cell.detailTextLabel?.text = String(articles![indexPath.row].text.filter { !"\n\r".contains($0) })
        cell.detailTextLabel?.font = UIFont.boldSystemFont(ofSize: 9.0)
        cell.detailTextLabel?.textColor = UIColor.darkGray
        
        return cell
    }
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        let linkURL = articles![indexPath.row].url
        let urlStr = linkURL.addingPercentEncoding(withAllowedCharacters: .urlQueryAllowed)
        let url:URL = URL(string:urlStr!)!
        let urlRequest = NSURLRequest(url: url)
        webView.loadRequest(urlRequest as URLRequest)
        
    }
    
    func webViewDidStartLoad(_ webView: UIWebView) {
        dotsView.isHidden = false
        dotsView.startAnimating()
    }
    
    func webViewDidFinishLoad(_ webView: UIWebView) {
        dotsView.isHidden = true
        dotsView.stopAnimating()
        webView.isHidden = false
        goButton.isHidden = false
        backButton.isHidden = false
        cancelButton.isHidden = false
        
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
    
    
    /*
     // MARK: - Navigation
     
     // In a storyboard-based application, you will often want to do a little preparation before navigation
     override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
     // Get the new view controller u
     sing segue.destinationViewController.
     // Pass the selected object to the new view controller.
     }
     */
    
}

