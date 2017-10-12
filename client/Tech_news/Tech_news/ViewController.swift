
import UIKit

var address = "http://133.130.91.14:8080"

class ViewController: UIViewController {
    
    
    
    var pageMenu:CAPSPageMenu?
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        var controllerArray:[UIViewController] = []
        
        let allViewController:UIViewController = AllViewController()
        allViewController.title = "Top"
        controllerArray.append(allViewController)
        let cookViewController:UIViewController = CookViewController()
        cookViewController.title = "CookPad"
        controllerArray.append(cookViewController)
        let eurekaViewController:UIViewController = EurekaViewController()
        eurekaViewController.title = "eureka"
        controllerArray.append(eurekaViewController)
        let voyageViewController:UIViewController = VoyageViewController()
        voyageViewController.title = "Voyage"
        controllerArray.append(voyageViewController)
        let ajitoViewController:UIViewController = AjitoViewController()
        ajitoViewController.title = "Ajito"
        controllerArray.append(ajitoViewController)
        let smartViewController:UIViewController = SmartViewController()
        smartViewController.title = "SmartNews"
        controllerArray.append(smartViewController)
        let wantedlyViewController:UIViewController = WantedlyViewController()
        wantedlyViewController.title = "Wantedly"
        controllerArray.append(wantedlyViewController)
        let cyberViewController:UIViewController = CyberViewController()
        cyberViewController.title = "Cyber Agent"
        controllerArray.append(cyberViewController)
        
        let paramerters:[CAPSPageMenuOption]=[
            .menuItemSeparatorWidth(5.0),
            .useMenuLikeSegmentedControl(false),
            .titleTextSizeBasedOnMenuItemWidth(true),
            .menuItemWidthBasedOnTitleTextWidth(true),
            .menuItemSeparatorPercentageHeight(0.1),
        ]
        
        pageMenu = CAPSPageMenu(viewControllers: controllerArray,
                                frame: CGRect(x:0.0, y:20.0, width:self.view.frame.width,height: self.view.frame.height), pageMenuOptions : paramerters)
        
        self.view.addSubview(pageMenu!.view)
        self.view.sendSubview(toBack: pageMenu!.view)
        
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }


}

