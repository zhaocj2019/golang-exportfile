<?php
/**
 *
 * User: changjiang
 * Date: 17/11/15
 * Time: 上午10:12
 */

namespace swoole\library;

use Yaf\Loader;
use \G7\Protocol\Http\Client;
use Yaf\Registry;
use \Exception;

use G7\Cache\Connector\RedisConnection;
use G7\Cache\Instance\Redis;
use swoole\action\exportData\data\ExportCustomDomainModel;
use G7\Alpha\Base\Context;

/**
 * 访问老框架的连接对象。
 * Class ConnectZen
 */
class ConnectZen
{

    public $domain;//访问的产品连接地址
    public static $envOldKey;//连接老框架使用的签名KEY
    public $link;//连接地址。
    public $linkParams;
    const LOG_CLASS = 'swoole_server';
    protected $exportId;
    /**
     * @var G7\Cache\Instance\Redis;
     */
    protected $cache;

    /**
     * ConnectZen constructor.
     *
     * @param $params
     *          appName string 应用名
     *          orgroot string
     *          orgcode string
     */
    public function __construct($params)
    {

        $this->linkParams = $params;

    }

    public function before()
    {
        $this->initConfig();
        return $this;
    }

    public function setExportId($exportId)
    {
        $this->exportId = $exportId;
        return $this;
    }

    /**
     * @throws Exception
     */
    protected function initConfig()
    {

        //初始化缓存内容
        $this->initRedis();

        //获得并缓存链接地址。
        $baseUrl = $this->getBaseUrlByAppName($this->linkParams);

        if (empty($baseUrl)) {
            throw new Exception("没有获取到业务方的链接地址！", 104);
        }
        $this->_initDomain($baseUrl);

        //连接地址。
        $this->_initBaseLink();
    }

    /**
     * 初始化Redis连接
     */
    protected function initRedis()
    {
        //设置缓存地址
        $redisConnection = RedisConnection::getInstance();
        $this->cache     = new Redis($redisConnection);

    }

    /**
     * @param $key
     * @param $value
     * @param int $expire
     */
    private function setCache($key, $value, int $expire)
    {
        if (!empty($value)) {
            $value = json_encode($value);
            $this->cache->set($key, $value);
            $this->cache->expire($key, $expire);
        }

    }

    /**
     * @param string $cacheKey
     * @param bool $isAsArray
     * @return  mixed|boolean
     */
    private function getCache($cacheKey, $isAsArray = false)
    {
        $content = $this->cache->get($cacheKey);

        if (!empty($content)) {
            return json_decode($content, $isAsArray);
        }
        return false;
    }

    /**
     * 根据应用名获得用户中心的配置信息
     * @param $arguments
     *                     appName  string 应用名称
     *                     orgroot string
     *                     orgcode
     * @return array
     * @throws Exception
     */
    public function getAppConfig($arguments)
    {

        if (empty($arguments['appName'])) {
            throw new \Exception("appName 为空!", 101);
        }
        //访问老框架的key ,此处暂未实现根据产品名称获得产品连接地址。
        //$envOldkey = self::getEnvOldKey();

        $argv            = [];
        $argv['orgroot'] = $arguments['orgroot'];
        $argv['orgcode'] = $arguments['orgcode'];

        $cacheKey = 'domain_' . md5(json_encode($argv));
        $res      = $this->getCache($cacheKey, true);

        if (empty($res)) {
            try {
                $client = new Client(['vega' => true]);

                $result = $client->request('POST', 'ucenter-v1.subsystem.getSubsystems',
                    [
                        'form_params' => $argv,
                        'timeout'     => '15'
                    ]
                )->getBody()->getContents();
                if (!empty($result)) {
                    $res = json_decode($result, true);
                }
                //缓存数据,180秒（半小时）
                $this->setCache($cacheKey, $res, 180);
                //记录本次请求的域名连接信息
//                Context::getLog()->log(G7_LOG_INFO, "EXPORT_ID:{$this->exportId}," . '访问的业务方的连接地址信息:' . var_export($res, true) . PHP_EOL . 'file:' . __FILE__ . 'line:' . __LINE__);

            } catch (Exception $e) {
                throw $e;
            }
        }
        foreach ((array)$res['data'] as $item) {
            if (!empty($item['code']) && $item['code'] === $arguments['appName']) {
                //记录用户访问的域名
                if ($item['code'] === 'project') {
                    $exportCustomData             = [];
                    $exportCustomData['domain']   = isset($item['baseurl']) ? $item['baseurl'] : '';
                    $exportCustomData['app_name'] = $item['code'];
                    (new ExportCustomDomainModel())->replaceData($exportCustomData);
                }
                return $item;
            }
        }

        return [];
    }

    private static $baseMessageCache;

    /**
     * 根据应用名获得用户中心的配置baseurl
     * @param $arguments
     *                     appName  string 应用名称
     *                     orgroot string
     *                     orgcode
     * @return mixed
     */
    public function getBaseUrlByAppName($arguments)
    {
        $config = $this->getAppConfig($arguments);
        //
        return isset($config['baseurl']) ? $config['baseurl'] : '';
    }

    /**
     * 初始化连接老框架的key
     * @return mixed
     */
    public static function getEnvOldKey()
    {
        if (self::$envOldKey === null) {
            self::$envOldKey = Registry::get('config')->get('extra')->get('envold')->get('key');
        }
        return self::$envOldKey;

    }

    /**
     * 初始化要访问的连接地址。
     * @param $domain
     */
    private function _initDomain($domain)
    {
        $this->domain = $domain;

    }

    /**
     * 连接地址
     */
    protected function _initBaseLink()
    {
        //  $this->link = $this->domain;
        $this->link = $this->domain . '/inside.php?t=json&m=index&f=service';

    }

}

/**
 * 访问老框架的操作类
 * Class Base_swoole_zendao_ConnectZenDao
 * @author karl.zhao<zhaochangjiang@huoyunren.com>
 */
class ConnectZenDao
{

    /**
     * @var Base_swoole_zendao_ConnectZenDao
     */
    protected static $thisObj;
    protected $params;
    protected $header = false;
    protected $execute_timeout = 120;
    protected $connect_timeout = 20;
    protected $exportId = '';
    protected $errorCode = null;

    protected function __construct()
    {
        $this->errorCode = SwooleErrorCode::singleton();
    }

    /**
     * @param $exportId
     * @return $this
     */
    public function setExportId($exportId)
    {
        $this->exportId = $exportId;
        return $this;
    }

    /**
     * @return Base_swoole_zendao_ConnectZenDao|ConnectZenDao|ConnectEws
     */
    public static function getInstance()
    {
        if (empty(self::$thisObj)) {
            self::$thisObj = new self();
        }
        return self::$thisObj;
    }

    protected function _orgParams()
    {


        $time                              = time();
        $dateTimeString                    = date('Y-m-d H:i:s', $time);
        $sendParams                        = $this->params['params']['remoteParams'];//传递的参数
        $customUser                        = [];
        $customUser['id']                  = '--mock-user-id--';
        $customUser['username']            = 'mockuser';
        $customUser['realname']            = 'mockuser';
        $customUser['roleid']              = '';
        $customUser['organ']               = [];
        $customUser['organ']['orgroot']    = isset($this->params['userInfo']['orgroot']) ? $this->params['userInfo']['orgroot'] : '';
        $customUser['organ']['orgcode']    = isset($this->params['userInfo']['orgcode']) ? $this->params['userInfo']['orgcode'] : '';
        $customUser['organ']['customerId'] = isset($this->params['userInfo']['customerId']) ? $this->params['userInfo']['customerId'] : '';
        $customUser['organ']['name']       = isset($this->params['userInfo']['name']) ? $this->params['userInfo']['name'] : '';
        $customUser['organ']['theme']      = isset($this->params['userInfo']['theme']) ? $this->params['userInfo']['theme'] : 'default';
        $params                            = [
            'customparams' => json_encode($sendParams),
            'customuser'   => json_encode($customUser),
            'method'       => "{$this->params['module']}.{$this->params['method']}",
            'format'       => 'json',
            'timestamp'    => $dateTimeString,

        ];
        $params['sign']                    = $this->sign($params);
        return $params;
    }


    /**
     * 访问老框架的签名算法。
     * @param $paramArr
     * @return mixed|string
     */
    function sign($paramArr)
    {
        //获得老框架的签名KEY
        $appSec = ConnectZen::getEnvOldKey();
        if (empty($appSec)) {
            $appSec = 'cat';
        }
        $sign = $appSec;
        ksort($paramArr);
        foreach ($paramArr as $key => $val) {
            if ($key != '' && !is_null($val)) {
                $sign .= $key . $val;
            }
        }
        $sign .= $appSec;
        $sign = strtoupper(md5($sign));
        return $sign;
    }

    /**
     * 获得本次请求的连接地址和参数
     * @return array
     */
    public function getRequestMessage()
    {
        return ['url' => $this->uri, $this->args];
    }

    /**
     * 访问老框架请求
     * @return array
     * @throws Exception
     */
    function sendRequest()
    {

        $requestUri = $this->uri;
        $args       = $this->args;
        if (empty($args)) throw new Exception('请求参数不能为空!', -90006);
        if (NULL == $requestUri) throw new Exception('请求地址不能为空!', -90007);
        $handle = curl_init();
        $error  = [];
        curl_setopt($handle, CURLOPT_URL, $requestUri);
        curl_setopt($handle, CURLOPT_POST, true);
        curl_setopt($handle, CURLOPT_USERAGENT, 'G7XClient');
        curl_setopt($handle, CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt($handle, CURLOPT_SSL_VERIFYHOST, false);
        curl_setopt($handle, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($handle, CURLOPT_POST, true);
        curl_setopt($handle, CURLOPT_POSTFIELDS, http_build_query($args));
//        print_r(http_build_query($args));
        curl_setopt($handle, CURLOPT_HEADER, $this->header);
        curl_setopt($handle, CURLOPT_NOSIGNAL, 1);
        curl_setopt($handle, CURLOPT_CONNECTTIMEOUT, $this->connect_timeout);
        curl_setopt($handle, CURLOPT_TIMEOUT, $this->execute_timeout);
        $data               = @curl_exec($handle);
        $code               = curl_getinfo($handle, CURLINFO_HTTP_CODE);
        $error['errorInfo'] = curl_error($handle);
        $error['errorNo']   = curl_errno($handle);
        curl_close($handle);
        return ['code' => $code, 'data' => $data, 'error' => $error];
    }

    /**
     * 获取微秒数
     *
     * @param string $mircrotime 微妙时间，默认为null则获取当前时间
     * @param boolean $get_as_float 获取微妙时间是否以浮点数返回,默认为false即不以浮点数方式返回
     * @return int
     */
    public function getMicroTime($mircrotime = null, $get_as_float = false)
    {
        return array_sum(explode(' ', $mircrotime ? $mircrotime : microtime($get_as_float)));
    }

    public function setParams($params)
    {

        $this->params = $params;

    }

    /**
     * 开始访问老框架
     * @param $params
     * @return array
     */
    public function run($params)
    {

        $this->setParams($params);

        $connect    = $this->connect();
        $params     = $this->_orgParams();
        $this->uri  = $connect->link;
        $this->args = $params;
        return $this->sendRequest();
    }

    protected $uri; //本次请求的地址
    protected $args;//本次请求的参数

    /**
     * 获得连接信息
     * //各产品名称
     * @throws  Exception
     * @return ConnectZen
     */
    public function connect()
    {

        if (isset($this->params['connectConfig'])) {
            $linkConfig = $this->params['connectConfig'];
        } else {
            $this->errorCode->setConst(SwooleErrorCode::PARAMS_WRONG);
            throw new Exception($this->errorCode->getMessage('缺少参数或参数为空'), $this->errorCode->getCode());
        }
        return (new  ConnectZen($linkConfig))->before()->setExportId($this->exportId);
    }


}