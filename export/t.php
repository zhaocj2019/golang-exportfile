<?php
/**
 * 连通Alpha框架获得信息
 * User: changjiang
 * Date: 17/11/13
 * Time: 下午2:04
 */
require_once dirname(__FILE__) . '/../../vendor/autoload.php';

//require_once __DIR__ . '/../vendor/autoload.php';
class AlphaFrameworkConnect
{
    private $fixer;

    /**
     * 连接zuul获取信息
     * @param $api
     * @param $args
     * @return array
     */
    public static function clientMessage($api, $args)
    {
        if(!isset($args['token'])&&isset($_COOKIE['_TOKEN'])){
            $args['token']=$_COOKIE['_TOKEN'];
        }
        $method = 'POST';
//        $args = ['username'=>'g7test'];
        $client     = new \G7\Protocol\Http\Client(['zuul' => true]);
        $response = $client->request($method,
            $api,
             [
                'form_params' => $args,//传递的参数
                'timeout'=>'30'
            ]
        )->getBody()->getContents();

        return json_decode($response);
    }

    
    

    /**
     * 创建新框架链接
     * @param $fixer
     * @return  array
     */
    public function start($fixer)
    {
        $this->fixer = $fixer;
        $api = $this->getApiByMethod('start');
//        $api ='newgatewayv1changjiang-v1.notice.fetchlist';
        $res = self::clientMessage($api, $this->fixer);
        return $res;
    }
    
    private static $exportConfig;
   
    /**
     * 
     * @return array
     */
    public  static function getConfig(){
       
        if(empty(self::$exportConfig)){
            self::$exportConfig= require dirname(__FILE__).'/../../config/Zuul.config.php';
        }
        return self::$exportConfig;
    }
    private function getApiByMethod($method)
    {
        $config = self::getConfig();
        return  $config->exportAppName .'.' ."exportdata." . $method;

    }

    /**
     * 下载地址
     * @param $fixer
     * @return array
     */
    public function download($fixer)
    {
        $this->fixer = $fixer;
        $api = $this->getApiByMethod('download');
        $res = self::clientMessage($api, $this->fixer);
        return $res;
    }


    /**
     * 获得下载列表
     * @param $fixer
     * @return array
     */
    public function getExportData($fixer)
    {
        $this->fixer = $fixer;
        $api = $this->getApiByMethod('getExportData');
        $res = self::clientMessage($api, $this->fixer);
        return $res;
    }

    /**
     * @param $fixer
     * @return  array
     */
    public function downloadCancel($fixer)
    {
        $this->fixer = $fixer;
        $api = $this->getApiByMethod('downloadCancel');
        $res = self::clientMessage($api, $this->fixer);
        return $res;
    }

    /**
     * @param $fixer
     * @return  array
     */
    public function getDownloadProgressByIds($fixer)
    {
        $this->fixer = $fixer;
        $api = $this->getApiByMethod('getDownloadProgressByIds');
        $res = self::clientMessage($api, $this->fixer);
        return $res;

    }
}
