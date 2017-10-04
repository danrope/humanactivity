package btree

//
//#include <string.h> 
//
//#define STRING_EQUAL(a,b) (_stricmp(a,b)==0)
//#define STRING_NOT_EQUAL(a,b) (_stricmp(a,b)!=0)
//#define STRING_SET(a,b)   strcpy(a,b)
//
//
//void _BTrees_Prediction_T14_16_28(
//      const double * _x__,
//      const double * _y__,
//      const double * _z__,
//      const double * _x1__,
//      const double * _x2__,
//      const double * _x3__,
//      const double * _x4__,
//      const double * _x5__,
//      const double * _x6__,
//      const double * _x7__,
//      const double * _x8__,
//      const double * _x9__,
//      const double * _x10__,
//      const double * _y1__,
//      const double * _y2__,
//      const double * _y3__,
//      const double * _y4__,
//      const double * _y5__,
//      const double * _y6__,
//      const double * _y7__,
//      const double * _y8__,
//      const double * _y9__,
//      const double * _y10__,
//      const double * _z1__,
//      const double * _z2__,
//      const double * _z3__,
//      const double * _z4__,
//      const double * _z5__,
//      const double * _z6__,
//      const double * _z7__,
//      const double * _z8__,
//      const double * _z9__,
//      const double * _z10__,
//      char * _pRet
//   )
//   {
//     double _MaxValue;
//     int _i;
//     double _den;
//     double _LearningRate;
//     double _PredictProb[3];
//     _MaxValue = -1.0E30;
//     _den = 0;
//     _LearningRate = 0.010000;
//
//     for(_i=0;_i<3;_i++)
//     {
//         _PredictProb[_i] = 0;
//     }
//
//    if ( _z3__ != NULL && *_z3__ <= 5.34300150000000e+000 ) {
//        _PredictProb[0]  += 1.000000 * 1.80519446430614e+000;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 5.34300150000000e+000 ) {
//        _PredictProb[0]  += 1.000000 * -2.59116674595316e-001;
//
//    }
//    else {
//    _PredictProb[0]  += 1.000000 * -1.79010964421550e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.25340425000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.15428097206713e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.25340425000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.08341004125783e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.20926778214486e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27352190000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.17260672997204e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27352190000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.04453370652040e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.96857258623446e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.26059115000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.26272801942311e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.26059115000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.01332676064948e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.90415965928831e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.23492370000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.23913967898921e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.23492370000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.89651816666686e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.78888599155291e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27710735000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.08226224680669e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27710735000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.90674338612511e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.92426136300678e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.28462560000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.07062004854749e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.28462560000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.91801904158459e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.53886208830868e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.25053490000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.11459142519261e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.25053490000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.82646212896089e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.64220755766751e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27433635000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.96379608994403e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27433635000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.80524742429849e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.42384381916812e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.28454875000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.87729051638107e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.28454875000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.80239279676785e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.50977111435645e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.28619700000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.96569818562926e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.28619700000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.83080914428099e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.69831823007686e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.26916340000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.93583616154418e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.26916340000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.73164014011497e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.12742824074489e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.25625235000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.01832250560383e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.25625235000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.70813022867803e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.41644816651504e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.25029415000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.00065019017589e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.25029415000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.65965638840859e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.93272153016371e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.29479990000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.84668926127065e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.29479990000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.70208364117939e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.02478524624216e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.25053490000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.02246914760216e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.25053490000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.60578942785232e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.04572220762009e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.24573710000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.02262116082897e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.24573710000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.58411691647358e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.70907241845369e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.27719280000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.85715365438049e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.27719280000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.56781129661911e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.14054336618642e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.31820015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.61423466177162e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.31820015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.58574009040385e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.20621205117838e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27710735000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.90854719509587e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27710735000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.55289968566416e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.44168994103374e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27925895000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.71445210035299e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27925895000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.52897075485586e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.80166984697471e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.27352190000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.80237933084876e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.27352190000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.47904824425714e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.07681702455762e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.31327920000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.63881828487216e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.31327920000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.53591077574550e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.41666249796908e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.27230625000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.68490926203398e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.27230625000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.45991837944103e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.48578097873713e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.25045155000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.81665778328496e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.25045155000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.41646154838251e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.35799185835896e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.31268425000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.59737414313372e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.31268425000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.49208162949089e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.19336223285314e-003;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.29611910000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.61931431668937e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.29611910000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.45514303215979e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.60166005903403e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.25637355000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.61820165933780e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.25637355000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.33886376116232e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.50647772662619e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.31820015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.48632055060383e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.31820015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.39046793981161e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.01537208708782e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.23911035000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.73133429526258e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.23911035000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.31080741551056e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.90779070942791e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.31911755000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.38404616906495e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.31911755000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.36158577314522e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.52166640709594e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.31117995000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.35935129136725e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.31117995000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.29909884343916e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.29994581718984e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20791065000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.77182184901150e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20791065000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.24206142848802e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.37336489140144e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.28482050000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.54819409247641e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.28482050000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.29319057724108e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.02293226576811e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.25042045000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.71135455144422e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.25042045000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.21331090589241e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.97875836377091e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.32016965000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.40813456997612e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.32016965000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.32610305927229e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.07123233342614e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.25053490000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.49353410700898e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.25053490000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.21329090160136e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.68069398366937e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.27439210000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.40713976042992e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.27439210000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.22717535643892e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.52820641060389e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.28457160000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.37724221967579e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.28457160000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.20002876079842e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.35035327376554e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23985575000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.54034820722005e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23985575000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.12815831112330e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.64833637017500e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23104310000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.43649085222782e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23104310000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.13035670447155e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.20964628103252e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20792805000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.62417307936340e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20792805000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.10470749846813e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.48655306061225e-004;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27705025000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.36178969105705e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27705025000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.20281925902882e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.56288789077518e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.31225315000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.28093578981417e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.31225315000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.19389628519035e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.35996575911642e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.28462560000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.27238156837041e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.28462560000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.14570237259997e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.03328116226626e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27230625000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.28159847771523e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27230625000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.11795589163030e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.31875470742795e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23911035000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.41710054961846e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23911035000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.07957676609451e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.57065867149743e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27346480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.26946960094273e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27346480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.09546359489178e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.04873452482704e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20572450000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.41306673650551e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20572450000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.00606442726386e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.03942842529018e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.31345120000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.15858081162973e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.31345120000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.10339278628162e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.48926595014805e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27346480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.15442875687833e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27346480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.05683840345136e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.57061487553744e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.31820015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.10181944760319e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.31820015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.10692117214793e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.53810918876012e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.32306365000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.06987459903158e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.32306365000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.07972495348144e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.86818070883735e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.27281270000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.06310529788200e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.27281270000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.01103018966375e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.54396029177991e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23997060000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.16456530234899e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23997060000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.77519747050324e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.23582715216488e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27346480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.14377532237937e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27346480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.00972512078557e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.06035069783832e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23738955000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.18472116512057e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23738955000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.61296096653661e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.40845272916063e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.36706155000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.90088459170816e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.36706155000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.07408556917325e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.50484528410205e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.32533125000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.91651600293816e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.32533125000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.03687527639675e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.67610775460582e-003;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.24358410000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.12605340207931e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.24358410000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.58874658812091e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.51606010776537e-004;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.29416910000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.01085131187167e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.29416910000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.01870106185659e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -7.12047520484042e-003;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.27433635000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.06085616560940e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.27433635000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.57574102332698e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.68094259912461e-003;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.28462560000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.00920134314028e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.28462560000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.76488566119046e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.53161427542936e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.31345120000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.86050712681951e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.31345120000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.56150598109471e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.43469007983710e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.24358395000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.06090863224611e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.24358395000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.21276496679401e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.65139940267523e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27346480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.83530082910727e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27346480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.20740265060140e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.83259767400316e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.28365730000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.84513718030354e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.28365730000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.46922491477801e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.14879608272041e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.32835145000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.80361004592995e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.32835145000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.78337152894095e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.68762393487531e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.32842840000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.73985871769649e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.32842840000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.76010451967336e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.56846833014423e-004;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20183090000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.05386086294641e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20183090000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.47021533091144e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.60137414870974e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.32812220000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.74935741783031e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.32812220000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.56093357834073e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.48189901150038e-003;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.28365730000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.84731514693305e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.28365730000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.20644638086948e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.51247677693419e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.32014990000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.61959939484375e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.32014990000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.97685132740876e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.35094172848557e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23343235000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.83167840715878e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23343235000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.55295985881451e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.43516618101918e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20292070000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.05545002191542e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20292070000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.12489949360027e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.63516632279647e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.28462560000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.76496912253255e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.28462560000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.89598706810510e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.05212337469812e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23492370000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.64951247202915e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23492370000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.29892464772160e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.87501550633695e-003;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.27401865000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.79813808804160e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.27401865000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.43360777130584e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.40563941117237e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.40268090000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.43751210001718e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.40268090000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.71470899394473e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.69135488225732e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.33362805000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.65411467505967e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.33362805000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.87876345007573e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.41646513839121e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23110055000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.72053527410165e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23110055000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.12940767925630e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -7.65171781569538e-003;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.31117995000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.61568775023558e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.31117995000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.05623041285115e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.46838993397593e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20775965000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.78065927424480e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20775965000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.51613158356770e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.27837045473832e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20787100000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.84399673728947e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20787100000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.48855317440156e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -5.46957863150910e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.21657045000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.75218732672488e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.21657045000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.72887207866078e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.34892794856045e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.24613355000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.73788629013081e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.24613355000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.16582769218975e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.45592357476206e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.32835145000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.50129460613505e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.32835145000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.67633758575990e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.17288605821885e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20555725000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.78478520994307e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20555725000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.50874831364707e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.43033327865485e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20775965000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.69179144062432e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20775965000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.41028800824801e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.66155530054437e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.29327945000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.41675360000691e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.29327945000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.16255429351663e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.68507595782245e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.29282215000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.45834733877104e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.29282215000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.11716839467993e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.92971792743119e-003;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.25637355000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.59004767595371e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.25637355000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.56977821065031e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.15113705949034e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.27489875000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.53084757319935e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.27489875000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.67762846165562e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.37129837688224e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.20802515000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.54302054146802e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.20802515000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.98489730496165e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.35512644116707e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.31820015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.50006659220332e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.31820015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.91044341400683e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.12329738452850e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.38008375000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.24632365241285e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.38008375000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.65483133723030e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.13961772944620e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.26107625000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.43819647582084e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.26107625000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.76672000125984e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.71387013203718e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.36302205000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.30378029786518e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.36302205000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.14855739706144e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.41087995177355e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.36391015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.37312312092150e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.36391015000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.54762785753384e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.23301880049630e-003;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 2.24328650000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.23214775517375e-002;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 2.24328650000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.16603492602422e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.20656204119694e-004;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 1.24455850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.44896971953131e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 1.24455850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 8.31447470018177e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.53165951278795e-002;
//
//    }
//    if ( _x9__ != NULL && *_x9__ <= 2.20628550000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -8.67986872346345e-002;
//
//    }
//    else if ( _x9__ != NULL && *_x9__ > 2.20628550000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.11048962603817e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.52278801028029e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.32533125000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.40198423007202e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.32533125000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.96893142112966e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.96372841259854e-003;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 1.28473850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.38765684694750e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 1.28473850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 8.43034271663132e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.10333512696590e-003;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.25569050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.05383823065651e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.25569050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.15570022234249e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.56435584818357e-003;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.22978350000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.27288068804422e-002;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.22978350000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.12474744817834e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.02318862495472e-003;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 6.95691150000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.46759974606475e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 6.95691150000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.69305494739929e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.91485219766884e-004;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 1.74107200000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.08454140718177e-001;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 1.74107200000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.36921629793609e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.25354509479246e-002;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.20808900000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.09956369172558e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.20808900000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.11953401577084e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.48800015570703e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.23104310000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.52925416498164e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.23104310000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.93957941832052e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.36015199232765e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.36494265000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.20032003780636e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.36494265000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.09011194299373e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.13888999904054e-002;
//
//    }
//    if ( _x9__ != NULL && *_x9__ <= 2.27174950000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.29643265162647e-002;
//
//    }
//    else if ( _x9__ != NULL && *_x9__ > 2.27174950000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.05568076196414e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.26060183777658e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23110055000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.57012140233133e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23110055000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.73455486219962e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.26467607689886e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.32055155000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.35333548396017e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.32055155000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.63274004758308e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.39329653185196e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 6.27226500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 7.19959421808734e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 6.27226500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.62603408722092e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.03600225334087e-002;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.25569050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.72218249356809e-002;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.25569050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.08291438213555e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.90144008856124e-003;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 1.28305100000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.34468682733434e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 1.28305100000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 7.73755858290556e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.37832758264753e-003;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 1.24570600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.41506726574395e-001;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 1.24570600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 7.53552224946519e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.44617628075987e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 1.22821310000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.50090098306626e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 1.22821310000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.72713989070029e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.28283574288956e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 6.91848600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.59305911864205e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 6.91848600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.58999064589603e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.44834703298497e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 6.25863150000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 7.20865928317375e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 6.25863150000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.41648846139940e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.20557919613765e-003;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.20980950000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.00799708130543e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.20980950000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.07803484212788e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.16359587448547e-002;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 1.65463250000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.14935628021200e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 1.65463250000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.34510297798501e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.05637423511848e-002;
//
//    }
//    if ( _x1__ != NULL && *_x1__ <= 2.27174950000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -8.70079365457877e-002;
//
//    }
//    else if ( _x1__ != NULL && *_x1__ > 2.27174950000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.05195924976103e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.65215540892529e-002;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 1.66852850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.03664925640796e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 1.66852850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 8.92296709944038e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.46054325222439e-002;
//
//    }
//    if ( _x7__ != NULL && *_x7__ <= 2.20037000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.24423446370638e-002;
//
//    }
//    else if ( _x7__ != NULL && *_x7__ > 2.20037000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.05411643622590e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -7.96795335153658e-003;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.25855850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.00957760181029e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.25855850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.04514171438398e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.86415712295878e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 6.95117600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.51581631220923e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 6.95117600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.52309776229315e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.30167109007740e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 8.38385500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.39466586689505e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 8.38385500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.86922476104118e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.87707227617095e-003;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 6.26263000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 7.27230687360919e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 6.26263000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.44775574639046e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -6.01780538938923e-003;
//
//    }
//    if ( _x9__ != NULL && *_x9__ <= 2.25569050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -8.74680694162939e-002;
//
//    }
//    else if ( _x9__ != NULL && *_x9__ > 2.25569050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.04850567481086e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.37748322717154e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.35536165000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.24923463387230e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.35536165000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.78120748159979e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.69131129334887e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.32835145000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.22795855640338e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.32835145000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.66340461228012e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.07357588418562e-003;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 6.51241000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.79951577312300e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 6.51241000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.45427994761506e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.53780483804033e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 8.23301600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.41028857361481e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 8.23301600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.83484966787843e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.23057823069342e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 8.31380050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.41676529246087e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 8.31380050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.98373788915117e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.79215205975059e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 6.34208750000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.87656953029905e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 6.34208750000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.56491720541438e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.25520822452833e-002;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 2.20808900000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.15796684317021e-002;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 2.20808900000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.05449214586047e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.64844613137847e-003;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 6.91217650000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.98095076455082e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 6.91217650000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.51167128820106e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.19223319071694e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.31911755000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.27084203667860e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.31911755000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.47894602699403e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.48501563831846e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 6.24975050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 7.12715487368697e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 6.24975050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.41848332324598e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.04817442831724e-002;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 2.20628550000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.37457585267304e-002;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 2.20628550000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.01895744354326e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.35711980725709e-002;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.20151700000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.33141313097113e-002;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.20151700000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.02083436915761e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.10683723355528e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 6.95440600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.23591981893859e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 6.95440600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.64941774925839e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.03893242119855e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.24346940000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.42898574651547e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.24346940000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.42574515168631e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.26999222114624e-003;
//
//    }
//    if ( _x9__ != NULL && *_x9__ <= 2.20808900000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.46057634878403e-002;
//
//    }
//    else if ( _x9__ != NULL && *_x9__ > 2.20808900000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.04416903092249e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.52704019042263e-003;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -7.54191500000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 1.05263825982821e+000;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -7.54191500000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -8.92141268488903e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -5.60576385586257e-003;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 1.36511480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.25780442433494e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 1.36511480000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.82076804110382e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.94610683153060e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.31830990000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.34986693907928e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.31830990000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.97685303198062e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.05230522888508e-004;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.36706155000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.20136475146905e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.36706155000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.86148025733078e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.78920201159961e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.20791065000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.49913799185494e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.20791065000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.08876299763850e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.09882052582259e-002;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.24422050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -8.63172755081947e-002;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.24422050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.73395256943653e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.01230471432086e-003;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -7.82867500000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 1.04069077593232e+000;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -7.82867500000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -9.61314881027294e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.25893810819673e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 6.34380850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.91395546029677e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 6.34380850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.50040830105282e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.01173412853196e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 6.92338600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.10548539983921e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 6.92338600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.65384765105090e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.04227779382713e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 8.38729550000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.09609860349908e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 8.38729550000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.84208337267498e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -5.11650317154288e-004;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 2.14331250000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.65799827394622e-002;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 2.14331250000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.66071340223527e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.57056085604195e-002;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 2.27117650000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -8.75406664403515e-002;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 2.27117650000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.81239747038238e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.03993221556538e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 6.75044150000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.03522388623952e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 6.75044150000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.56444025140396e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.88049556396157e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 6.27169100000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.72989233654482e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 6.27169100000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.45172918712850e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.48380804906445e-003;
//
//    }
//    if ( _x9__ != NULL && *_x9__ <= 2.20374600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -8.90015145843312e-002;
//
//    }
//    else if ( _x9__ != NULL && *_x9__ > 2.20374600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.06728225254228e+000;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.54835357661235e-003;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 2.25225000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.13062628706813e-002;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 2.25225000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.48135206777444e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.73313878965334e-003;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.33266915000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.13860429380601e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.33266915000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.19213689064897e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.78319701293927e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 6.62311850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.26660298341043e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 6.62311850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.44801158420185e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.10989700337962e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 8.34829500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.29590644349220e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 8.34829500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.96643010477698e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.05619025339193e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.36627970000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.08745185242770e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.36627970000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.46339474619803e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.61940943051441e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 6.95117600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.94182412347270e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 6.95117600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.56536298964323e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -5.22732635269929e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.36935860000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.23077084386963e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.36935860000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.23865112472466e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.03121426213848e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 8.34829250000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.85523925828783e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 8.34829250000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.83897353129679e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.35087358828388e-003;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.20151700000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.03766176112607e-001;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.20151700000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.76855742418571e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.92800638895140e-002;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 2.20776050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -8.64055083325180e-002;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 2.20776050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.24610347031371e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.21383676079646e-002;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -5.68367500000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 9.43753519000956e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -5.68367500000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -9.45578949754762e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.39650449920457e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.38523575000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.15148882144021e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.38523575000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.63260472533806e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -6.26563999772910e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 7.75870500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.25467803151734e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 7.75870500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.68047955687758e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.76739549263658e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 6.91317600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.84181377477874e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 6.91317600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.45733882824281e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.36261305364193e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.36494265000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.18888009163767e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.36494265000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.65344635838141e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.30494000706975e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 8.15501450000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.00716224164357e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 8.15501450000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.87209659912926e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.88589855222124e-003;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.19375050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.47682423026143e-002;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.19375050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.49975856569606e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.31376208082591e-003;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.20808900000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.01415006411959e-002;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.20808900000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.68824833410723e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.85106236355213e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.36494265000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.14267927646474e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.36494265000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.32311106320563e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.18010060188894e-002;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 1.23538250000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.24628003653557e-001;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 1.23538250000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.72130085535028e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.79913213190445e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.23370420000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.32660651909806e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.23370420000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.09078569770856e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.92399163780570e-003;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -5.72563000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 9.65077080352412e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -5.72563000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -9.76386496847078e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.35334632284265e-004;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.22300050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -9.40256265001644e-002;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.22300050000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.37484009383822e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.72963988454406e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23426820000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.31449376288882e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23426820000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.82462730799324e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.36790571028343e-002;
//
//    }
//    if ( _x2__ != NULL && *_x2__ <= 2.25465850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -8.66763332268476e-002;
//
//    }
//    else if ( _x2__ != NULL && *_x2__ > 2.25465850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.74528207809666e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -7.12279295364768e-003;
//
//    }
//    if ( _x__ != NULL && *_x__ <= 1.01740850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.40802630911918e-001;
//
//    }
//    else if ( _x__ != NULL && *_x__ > 1.01740850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.44010337082277e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.39481397267774e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 8.31491500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.02235461441951e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 8.31491500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.92386798659733e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.73421502243754e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.19317540000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.45411212125484e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.19317540000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.55179061058914e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.01435071043430e-002;
//
//    }
//    if ( _x10__ != NULL && *_x10__ <= 2.25523200000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -8.55814041601689e-002;
//
//    }
//    else if ( _x10__ != NULL && *_x10__ > 2.25523200000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 9.48624492217650e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.45841285727028e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 8.38786850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.58744471101169e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 8.38786850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.84951172634843e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.38353215175362e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 6.91848600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.52823312477119e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 6.91848600000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.43060713956101e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.02783153556838e-002;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= -4.84960000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 8.85560489744688e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > -4.84960000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -1.01010194403617e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.18595150909164e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.23104310000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.44817361334743e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.23104310000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.86829583219807e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.83833705618585e-003;
//
//    }
//    if ( _x9__ != NULL && *_x9__ <= 1.24575500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.32547274930458e-001;
//
//    }
//    else if ( _x9__ != NULL && *_x9__ > 1.24575500000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.52027830203152e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.55383958485902e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.40101765000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.13288169445825e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.40101765000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.70321360589361e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.52838976759057e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 8.38786850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.56410474145967e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 8.38786850000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.87724047297837e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.25799164421504e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.27273715000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.25679789800127e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.27273715000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.11973937945858e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.35087342898715e-002;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= -7.50652000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 9.84277499798619e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > -7.50652000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -8.95793712956353e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.01564396273452e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 8.33074350000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.53179860815744e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 8.33074350000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.79041495405521e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.11236402268708e-002;
//
//    }
//
//   if(_MaxValue <= _PredictProb[0])
//   {
//     _MaxValue = _PredictProb[0];
//     STRING_SET(_pRet,"Jogging");
//   }
//    if ( _z4__ != NULL && *_z4__ <= 1.44010290000000e+001 ) {
//        _PredictProb[1]  += 1.000000 * 1.10540211830390e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 1.44010290000000e+001 ) {
//        _PredictProb[1]  += 1.000000 * -9.41865040729629e-001;
//
//    }
//    else {
//    _PredictProb[1]  += 1.000000 * 5.41631629494823e-003;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00327500000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.62040605117525e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00327500000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.92355924631608e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.00238301418766e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 1.00388300000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.47828977188181e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 1.00388300000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.30462387478151e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.50768257992953e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 1.00298825000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.51662785338106e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 1.00298825000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.41856367910193e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 8.79018696349390e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00298825000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.51469908277381e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00298825000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.39089143588640e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.74237694265127e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00447965000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.58712991665982e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00447965000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.52606972222554e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.90021943604640e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 1.00447945000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.49704953653072e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 1.00447945000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.17498089624127e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.55331891290131e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00482385000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.56523931079628e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00482385000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.27567338241426e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.93161388455647e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00490060000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.32328374625201e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00490060000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.45591189625688e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.53691585116248e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00303420000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.38061635376147e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00303420000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.73315516660959e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.32604029939951e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00314885000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.33359279659148e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00314885000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.40008660049339e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.46550703597641e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 1.00308015000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.23486208116831e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 1.00308015000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.43182111118791e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.65934096909046e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00484330000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.33733543691458e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00484330000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.38697963183333e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.33111840368902e-002;
//
//    }
//    if ( _y1__ != NULL && *_y1__ <= 5.24517000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 6.50133502067198e-001;
//
//    }
//    else if ( _y1__ != NULL && *_y1__ > 5.24517000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -1.08367534345155e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.59186115268891e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 1.00450575000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.32960602742058e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 1.00450575000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.14385908046213e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.54689199091721e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00486305000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.33438799310503e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00486305000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.26759435635511e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.67870511543679e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00490060000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.32899458313002e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00490060000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.27139003665697e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.83495375603123e-002;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= 5.55045000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 6.27030715093806e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > 5.55045000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -1.08268878787455e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.78807053502679e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.24063160462487e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.07898774253749e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.06139463941771e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00348645000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.31034704040713e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00348645000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.10875888304206e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.85360425292470e-002;
//
//    }
//    if ( _y3__ != NULL && *_y3__ <= 5.19421000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 6.36974288287433e-001;
//
//    }
//    else if ( _y3__ != NULL && *_y3__ > 5.19421000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -1.06802721665368e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.83882868376656e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00473015000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.26604604980527e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00473015000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.82298204905505e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.25662226832167e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.26998745046859e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.71952695724889e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.44210275046233e-002;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= 5.27073500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.81610959054613e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > 5.27073500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -1.00846166868315e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.19332029879687e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00447965000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.14087244359637e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00447965000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.66780169912176e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.75577939026589e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00338980000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.20798722431687e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00338980000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.46746685532941e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.48061513346673e-002;
//
//    }
//    if ( _y3__ != NULL && *_y3__ <= 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.78723714462082e-001;
//
//    }
//    else if ( _y3__ != NULL && *_y3__ > 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.69320263906181e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.35420467881076e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 1.00356170000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.21260545651578e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 1.00356170000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.50627131673284e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.31210405551006e-002;
//
//    }
//    if ( _y1__ != NULL && *_y1__ <= 5.53455500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.58600255236424e-001;
//
//    }
//    else if ( _y1__ != NULL && *_y1__ > 5.53455500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.71490199581876e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.29243173149951e-002;
//
//    }
//    if ( _y3__ != NULL && *_y3__ <= 5.24517500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.33016357557300e-001;
//
//    }
//    else if ( _y3__ != NULL && *_y3__ > 5.24517500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.14245535669402e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.19353290742133e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.20376647449953e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.41876488512001e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.00855499538594e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.14242500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.18228042085181e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.14242500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.70524916433542e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.65075492462576e-002;
//
//    }
//    if ( _y2__ != NULL && *_y2__ <= 5.23960000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.18064755548028e-001;
//
//    }
//    else if ( _y2__ != NULL && *_y2__ > 5.23960000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.98296565642712e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 7.89229785576033e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00298825000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.18511085509552e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00298825000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.31425907180240e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.35746535204902e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48617600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.68794921766580e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48617600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.47510023942719e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.83984507544398e-002;
//
//    }
//    if ( _y2__ != NULL && *_y2__ <= 5.28220500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.51332455948076e-001;
//
//    }
//    else if ( _y2__ != NULL && *_y2__ > 5.28220500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -1.00520309912517e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.50604392424632e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00308015000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.42413120885282e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00308015000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.64421876703485e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.88329757198856e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 9.44029450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.64220358158282e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 9.44029450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.47291816494045e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.18172554725560e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.61215734190977e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.38452213892124e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.94473714093376e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.44647300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.31292237644543e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.44647300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.36702539650331e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.31415029200089e-002;
//
//    }
//    if ( _y3__ != NULL && *_y3__ <= 5.10392000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.20862201362525e-001;
//
//    }
//    else if ( _y3__ != NULL && *_y3__ > 5.10392000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.61557913046190e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.67681976821034e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.51255950000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.40403958255148e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.51255950000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.59208204845483e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.23924848945780e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48560350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.45172761575412e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48560350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.49478230013201e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.44665171034481e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00356170000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.23667796815850e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00356170000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.28059848361892e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.39189333295823e-002;
//
//    }
//    if ( _y3__ != NULL && *_y3__ <= 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.76106139486219e-001;
//
//    }
//    else if ( _y3__ != NULL && *_y3__ > 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.63243634679108e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.85336279625796e-002;
//
//    }
//    if ( _y5__ != NULL && *_y5__ <= 5.16045000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.99717084185931e-001;
//
//    }
//    else if ( _y5__ != NULL && *_y5__ > 5.16045000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.16394384720019e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.55417669055399e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.33619139882917e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.43417022016178e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.54818588311148e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46208850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.50487378870134e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46208850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.58213504494295e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.42751041382022e-002;
//
//    }
//    if ( _y1__ != NULL && *_y1__ <= 5.16045000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.74812542413754e-001;
//
//    }
//    else if ( _y1__ != NULL && *_y1__ > 5.16045000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.01037848981681e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.58654344516592e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.61321467846691e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.47654884734425e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.32999339076501e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.65180665544183e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.26286216555302e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.92988509296096e-003;
//
//    }
//    if ( _y3__ != NULL && *_y3__ <= 5.14914500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.70790832616425e-001;
//
//    }
//    else if ( _y3__ != NULL && *_y3__ > 5.14914500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.18560155996964e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.40896550185717e-002;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= 5.23960000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.03645743724194e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > 5.23960000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.75136644554710e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.11965943120055e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00128245000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.13913370491652e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00128245000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.87072060427782e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.25230369729385e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00320465000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.08323883831464e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00320465000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.77578884647719e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.62608071502473e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00338980000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.16654037745372e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00338980000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.81494237321965e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.96202645624000e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48535650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.94635789742705e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48535650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.42480897755773e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.08120276353645e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.24517000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.68608718588434e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.24517000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.92583632884587e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.35763153631707e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.26470306158257e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.49048270562775e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.87982504370337e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.06257900195611e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.43220418031278e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.14444356369794e-003;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.48491350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.24735751711946e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.48491350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.47675238214345e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.66198999941558e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46739550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.21551995468050e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46739550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.45387644319174e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.06922806922686e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.11113574016598e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.51908793711478e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.09878688328137e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.97551583847860e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.45424751100808e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.46871014312252e-002;
//
//    }
//    if ( _y__ != NULL && *_y__ <= 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.46623404095011e-001;
//
//    }
//    else if ( _y__ != NULL && *_y__ > 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.13345760000804e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.50935699732812e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.16176500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.37282459919797e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.16176500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.04704370672679e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.07001232727841e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48378300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.17343801301241e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48378300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.48755345078533e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.22895732976761e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00473015000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.11105858517772e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00473015000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.66782306330536e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.13018166229782e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 9.46682400000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.16910997828519e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 9.46682400000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.36851529890885e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.72125458043341e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00602800000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.08943253286582e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00602800000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.61935859551369e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.42578371117516e-002;
//
//    }
//    if ( _y5__ != NULL && *_y5__ <= 5.61256000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.39015585659571e-001;
//
//    }
//    else if ( _y5__ != NULL && *_y5__ > 5.61256000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.32904349612690e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.05115565622175e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00327500000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.04866763887322e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00327500000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.49686180578221e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 9.87892508034322e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48560350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.22176139988272e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48560350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.48746526830698e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.23101975167996e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.08382768168028e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.60701691524735e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.60912879028416e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48617600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.24301119673838e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48617600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.50568099125249e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.67811010780176e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.90023618102704e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.42097325380913e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.10795271440068e-002;
//
//    }
//    if ( _y3__ != NULL && *_y3__ <= 5.01838500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.28325289732302e-001;
//
//    }
//    else if ( _y3__ != NULL && *_y3__ > 5.01838500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.47956824915302e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.73567672071907e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 1.00297860000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.03433586248603e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 1.00297860000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.52323491907234e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.26494894192338e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00303420000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.05939879227695e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00303420000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.50561254310345e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.85533621459317e-003;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 9.45043650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.12784432603937e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 9.45043650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.32284795842166e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.57820688854036e-002;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= 5.16176500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 4.31919275071527e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > 5.16176500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.88307476273503e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.70333890369370e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.90172820071640e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.46558051551180e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.42182429242698e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46323450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.94950325641010e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46323450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.35931180752156e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 7.26604187882858e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46208850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.19642302963074e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46208850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.50773964168830e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.11294082134503e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46474350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.94376144160616e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46474350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.39905283050723e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.97142679647640e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48378300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.84451907024956e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48378300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.43819844426488e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.58600774348666e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.45061650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.90644617025430e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.45061650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.35358007411729e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.18381561634713e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48547200000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.84022908406812e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48547200000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.41564327684046e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.83724700144148e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.73371559632575e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.30750255032570e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.98165237282356e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.76681607168489e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.46151823693382e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.36975666595096e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46266050000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.86553772498781e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46266050000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.38372253326954e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.00797777606104e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.72899301999992e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.36048649902142e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.59803395202350e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.45383000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.63168600344387e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.45383000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.32120126599487e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.28088332660219e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46474350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.73569034953701e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46474350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.37951730541809e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.30361479192614e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 4.53534050000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.28077140003899e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 4.53534050000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.94091413477918e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.31507523432585e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 9.46117150000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.66656238265027e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 9.46117150000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.25630321969054e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.27813373420079e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 4.79020000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.94002102373117e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 4.79020000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 1.00184742507546e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.85394073036558e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46265950000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.87240542783695e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46265950000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.36410053600813e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.14859505617885e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.44660200000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.86081021470469e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.44660200000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.40616139240617e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.95949321990418e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48560350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.76525546470659e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48560350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.35906934819071e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.05302020758566e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 3.94301450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.55973947053973e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 3.94301450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.93210903221144e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.31573437130671e-002;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= 5.53455500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.86690177615607e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > 5.53455500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.60495746643020e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.59431372487790e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.57512636591923e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.29363433021445e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.89718702859567e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.51544300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.64340656292117e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.51544300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.46043729560673e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.24649135596592e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48547200000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.76393558130476e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48547200000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.32718304261130e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.21926511459954e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 4.62010750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.93069649205361e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 4.62010750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 9.16292978094990e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.66831143526626e-002;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.97871943466777e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.91526860260221e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.30231675723704e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46208850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.67489499194390e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46208850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.23504772992912e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.43875307682445e-002;
//
//    }
//    if ( _y3__ != NULL && *_y3__ <= 5.55045000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.90920851751587e-001;
//
//    }
//    else if ( _y3__ != NULL && *_y3__ > 5.55045000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.81276801645767e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.47801500573924e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 4.52916100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.06630335697726e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 4.52916100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.68751989921225e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.70324207206384e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.60528704285213e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.34192670067067e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.19113653253363e-003;
//
//    }
//    if ( _y3__ != NULL && *_y3__ <= 5.20765000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.72877548714370e-001;
//
//    }
//    else if ( _y3__ != NULL && *_y3__ > 5.20765000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.16653498111467e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.29469876666758e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.73671404386148e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.38138004087893e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.43593377332113e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 4.53501900000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.92911873008993e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 4.53501900000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.94070739737885e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.90027728424213e-002;
//
//    }
//    if ( _y__ != NULL && *_y__ <= 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.85321168143169e-001;
//
//    }
//    else if ( _y__ != NULL && *_y__ > 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.37123247845273e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 9.44957033604320e-004;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.45061650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.72325286782926e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.45061650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.31075633687360e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.69235623283676e-002;
//
//    }
//    if ( _y1__ != NULL && *_y1__ <= 5.64926500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.52850654338487e-001;
//
//    }
//    else if ( _y1__ != NULL && *_y1__ > 5.64926500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.52387929847049e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.04819357917602e-004;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 4.76832350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.83121980940041e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 4.76832350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.56482538819504e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.70682238365199e-002;
//
//    }
//    if ( _y4__ != NULL && *_y4__ <= -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.87681143626962e-001;
//
//    }
//    else if ( _y4__ != NULL && *_y4__ > -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.32008730352774e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.15420365832275e-002;
//
//    }
//    if ( _y1__ != NULL && *_y1__ <= 5.16176500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.66650668727402e-001;
//
//    }
//    else if ( _y1__ != NULL && *_y1__ > 5.16176500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.25788941350670e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 8.61864449636768e-003;
//
//    }
//    if ( _y__ != NULL && *_y__ <= 5.53914500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.68392275065002e-001;
//
//    }
//    else if ( _y__ != NULL && *_y__ > 5.53914500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.62993599387563e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.34669160267391e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 4.04975600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.42718249409279e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 4.04975600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.71132216725688e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.19579012120560e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 4.57876300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.38332725880912e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 4.57876300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 9.27375998934557e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.13852248800659e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.61060637703708e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48617500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.30783156700591e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.50995086632077e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.47305000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.71817889484471e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.47305000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.33014449236618e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.44278835759416e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 9.50526550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.53110899520543e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 9.50526550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.38626344382671e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.66038461542350e-002;
//
//    }
//    if ( _y1__ != NULL && *_y1__ <= 5.14242500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.91414694745304e-001;
//
//    }
//    else if ( _y1__ != NULL && *_y1__ > 5.14242500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.79889943189897e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.88758250876393e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 4.51949350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.10424785910518e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 4.51949350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.29811028186360e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.05455252746327e-002;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.90519415471339e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.43239960178586e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.16617998128699e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 4.53501900000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.93645030555810e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 4.53501900000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.78088335272713e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.12124136916675e-003;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00298825000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -9.33434420397457e-002;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00298825000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.12746068410904e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.40842690940498e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.51028421745095e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.30569844918729e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.55683350470151e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 4.10417600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.16577746702599e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 4.10417600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.83783937586855e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.15520980249885e-002;
//
//    }
//    if ( _y2__ != NULL && *_y2__ <= 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.80048292493118e-001;
//
//    }
//    else if ( _y2__ != NULL && *_y2__ > 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.96987302146468e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.81775052376602e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.51265650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.66895244822238e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.51265650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.44155051566815e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.63149761174093e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.80098675845325e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.20469697023469e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.66741457418116e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.43573800000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.51241469764976e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.43573800000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.11879506060961e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.82079970342553e-002;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= -1.19451450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.65906404533490e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > -1.19451450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.35325445742461e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.62575825714033e-002;
//
//    }
//    if ( _y8__ != NULL && *_y8__ <= -8.28750000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.77424949026744e-001;
//
//    }
//    else if ( _y8__ != NULL && *_y8__ > -8.28750000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 7.02814336264300e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.02965855134848e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.51655800000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.42964894556784e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.51655800000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.33373144452127e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.92254681372775e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 3.83002950000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.41031707817390e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 3.83002950000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.00356023664935e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.11066544815303e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -9.97924422487197e-002;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 1.00338985000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.25175162161008e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.65060561411002e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.51265650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.42746551770680e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.51265650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.21726897446934e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.23371914070265e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.45061650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.59539430186789e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.45061650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.27768013221145e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.38258584611839e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.43685000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.53847303078243e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.43685000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.26979720740903e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.27555685424211e-003;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.53455500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.51054748262444e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.53455500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.90907663984053e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.47943582458135e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 3.91868100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.56135550832448e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 3.91868100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.94686922870842e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.81960748884238e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 3.85297050000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.44356211568042e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 3.85297050000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.54515718843984e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.90552220044175e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.62639371734629e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.42483861432017e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.35762228139109e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 4.52855500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.70778325657498e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 4.52855500000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.57983348308974e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.66690331152212e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.61256000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.41844553447820e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.61256000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.28486801109975e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.39376712669496e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.50526550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.41516302913520e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.50526550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.34276004606722e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.06274593816373e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 4.78782350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.98377946043391e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 4.78782350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.31595838696655e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.87759702104747e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00303420000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -9.04251601144075e-002;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00303420000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.08841306709945e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.38863244722355e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46230000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.51616851277196e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46230000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.30119943487101e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.50174445409244e-006;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.64926500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.52171977954070e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.64926500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.69727971183137e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.10834579349226e-002;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -9.09731162945839e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.40560205887777e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.86868360976909e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.45980067273415e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.19976239173555e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.80639582248114e-002;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -1.19036850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.92482896717842e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -1.19036850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 5.50456696521108e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.27817939628682e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 4.58708800000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.77654936774643e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 4.58708800000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.80566246097528e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.17719025080838e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 3.53715300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.67941850015384e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 3.53715300000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.84466819339304e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.27789299597374e-003;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= 5.13784000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.37678071913191e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > 5.13784000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.36124846832462e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.36796889282874e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.51141050000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.35511700191732e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.51141050000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.19001745468486e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 8.50519875812690e-003;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00189850000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -9.16045850340089e-002;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00189850000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.80510955536784e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.26841162714544e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.42573376477939e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48491450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.16898380047314e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.07302139012391e-002;
//
//    }
//    if ( _y8__ != NULL && *_y8__ <= -1.25475100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -9.23539145188579e-001;
//
//    }
//    else if ( _y8__ != NULL && *_y8__ > -1.25475100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 5.86802810935612e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.63541895699617e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 3.83330700000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.35343499719293e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 3.83330700000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.45915064619204e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 9.67926962735093e-003;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= 5.26205500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.52529954615070e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > 5.26205500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.15328819343128e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 7.17267592151537e-003;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 9.51265650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.43588832718139e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 9.51265650000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.24963947523325e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.68611395459306e-004;
//
//    }
//    if ( _y8__ != NULL && *_y8__ <= 5.29040000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.54872770771729e-001;
//
//    }
//    else if ( _y8__ != NULL && *_y8__ > 5.29040000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.49389305258651e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.49135788609431e-003;
//
//    }
//    if ( _y__ != NULL && *_y__ <= -9.79522500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.51685413073482e-001;
//
//    }
//    else if ( _y__ != NULL && *_y__ > -9.79522500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 6.72533142287578e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.00456573920844e-003;
//
//    }
//    if ( _y8__ != NULL && *_y8__ <= 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.61784748727909e-001;
//
//    }
//    else if ( _y8__ != NULL && *_y8__ > 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.84066466328811e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.86282335532142e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.64926500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.29292209139253e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.64926500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.73748162131284e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.09879609230622e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.51302850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.41963036135212e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.51302850000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.17640916873288e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.77297929753628e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.14914500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.42240465718803e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.14914500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.03092375508973e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.60217811570474e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48547200000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.39724907569277e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48547200000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.19154992848598e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.08209370840664e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.62837428626702e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -6.79096716071871e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.18040017619347e-002;
//
//    }
//    if ( _y1__ != NULL && *_y1__ <= -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.62171429154030e-001;
//
//    }
//    else if ( _y1__ != NULL && *_y1__ > -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 5.65986405969448e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.20085974493798e-002;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -1.18089700000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.41251593205579e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -1.18089700000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 5.76941090919868e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.04749455357312e-002;
//
//    }
//    if ( _y6__ != NULL && *_y6__ <= -1.19149900000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.99114214043184e-001;
//
//    }
//    else if ( _y6__ != NULL && *_y6__ > -1.19149900000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.44096693349275e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.28240697454081e-002;
//
//    }
//    if ( _y2__ != NULL && *_y2__ <= 5.27073500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.45684171748490e-001;
//
//    }
//    else if ( _y2__ != NULL && *_y2__ > 5.27073500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.53838763300943e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.69677601420988e-002;
//
//    }
//    if ( _y1__ != NULL && *_y1__ <= 5.64926500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.51489039386781e-001;
//
//    }
//    else if ( _y1__ != NULL && *_y1__ > 5.64926500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.67655080470518e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.82428685727898e-002;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.96139781879869e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 5.42980392906179e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.87975197412413e-002;
//
//    }
//    if ( _y__ != NULL && *_y__ <= -9.75000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.85210398943584e-001;
//
//    }
//    else if ( _y__ != NULL && *_y__ > -9.75000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.78068853760948e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.73967828844404e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 4.04287400000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.10391827981834e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 4.04287400000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.57373142773883e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.62594803964202e-002;
//
//    }
//    if ( _y2__ != NULL && *_y2__ <= 5.26779000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.52108112551754e-001;
//
//    }
//    else if ( _y2__ != NULL && *_y2__ > 5.26779000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.61382412245119e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.27090512615697e-003;
//
//    }
//    if ( _z5__ != NULL && *_z5__ <= 3.93957400000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.49517434342590e-001;
//
//    }
//    else if ( _z5__ != NULL && *_z5__ > 3.93957400000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.78280208782450e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.80904694487887e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.48560350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.36411868932664e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.48560350000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.19555396657715e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.53692219257361e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 3.91868100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.48548303721485e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 3.91868100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.59972338842006e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 8.11757041821728e-004;
//
//    }
//    if ( _y6__ != NULL && *_y6__ <= -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.71202503335405e-001;
//
//    }
//    else if ( _y6__ != NULL && *_y6__ > -1.19092550000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.01384080328960e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.35297257893876e-002;
//
//    }
//    if ( _y7__ != NULL && *_y7__ <= 5.61485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.35801206018350e-001;
//
//    }
//    else if ( _y7__ != NULL && *_y7__ > 5.61485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.13739323480793e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.54683490679630e-004;
//
//    }
//    if ( _y__ != NULL && *_y__ <= 5.13784000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.40015614494214e-001;
//
//    }
//    else if ( _y__ != NULL && *_y__ > 5.13784000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.64262779754354e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.47745730744673e-003;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 9.48560450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.32024310948388e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 9.48560450000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.26522877952616e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.72536779075945e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 3.91868100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.46126837852597e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 3.91868100000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.54227620810677e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.60291560421406e-005;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= 5.13784000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.46813698612184e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > 5.13784000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.40920787049545e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.28057672258622e-002;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= -1.19149900000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -8.90241253840859e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > -1.19149900000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.20856887881448e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.03488214211427e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.25943996351707e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 9.46208750000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.24064590197851e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.94258897774920e-004;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 3.94917600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.28598004501156e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 3.94917600000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.65041444039854e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.56812519733308e-002;
//
//    }
//    if ( _y__ != NULL && *_y__ <= 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.48664095079154e-001;
//
//    }
//    else if ( _y__ != NULL && *_y__ > 5.22485500000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -7.26216825198407e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.95040062379165e-003;
//
//    }
//    if ( _y2__ != NULL && *_y2__ <= 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.53177234617530e-001;
//
//    }
//    else if ( _y2__ != NULL && *_y2__ > 5.27909000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -8.52553039343586e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.42385282887792e-002;
//
//    }
//
//   if(_MaxValue <= _PredictProb[1])
//   {
//     _MaxValue = _PredictProb[1];
//     STRING_SET(_pRet,"Walking");
//   }
//    if ( _z3__ != NULL && *_z3__ <= 9.26503950000000e+000 ) {
//        _PredictProb[2]  += 1.000000 * -9.69478994043369e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 9.26503950000000e+000 ) {
//        _PredictProb[2]  += 1.000000 * 4.45434428428119e-001;
//
//    }
//    else {
//    _PredictProb[2]  += 1.000000 * -2.02660314084382e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00671640000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.57936778721602e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00671640000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.32086021202182e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -9.98453966005099e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.58205038844245e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.33047333380529e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.90182718213682e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00447965000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.51260337351080e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00447965000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.30163070018650e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.48959795446883e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.46687298347308e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.30304603699311e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.79604718144966e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00574120000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.43213338511704e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00574120000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.25267271626671e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.11317285791748e-001;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.29380471330414e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.26922282044630e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.44136088587339e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.21799727074828e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.25656442323690e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -8.50555951637372e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00683110000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.14177021994121e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00683110000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.24165780788213e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.76954088770492e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00792070000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.13223322892409e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00792070000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.22784017301331e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.96050377318587e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.14209283270373e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.22991485673308e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.80232335888798e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00706030000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.04623398531191e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00706030000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.21743857106616e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.41620906182907e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00693240000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 5.01007033521200e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00693240000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.19256290274734e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.66263182605405e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.94668123579356e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.17807197727958e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.58062407004528e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.90727716409301e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.17145899012535e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.28606184426920e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00490060000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.81999326687495e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00490060000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.15444152166146e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.12563676364777e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00704710000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.92693236663915e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00704710000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.13762110188688e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.65571532314907e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00586575000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.72733336988174e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00586575000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.14419004554982e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.45594921119992e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.58945392843226e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.13025986802700e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.83438910318546e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.69209472046358e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.12390199780298e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -8.99955458296611e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00591995000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.58904289513893e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00591995000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.11672483690949e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -8.66714380683519e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.60706495300051e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.09713359084115e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.83264422703228e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.46827987293302e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.09704321443820e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.71411441577296e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.52015123895117e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.09640638166965e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -8.35124727106989e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00382550000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.43582138172248e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00382550000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.07978250909900e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.58242217068959e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00388295000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.46231811095252e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00388295000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.06729973519168e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.26810670187438e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.29314370023281e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.07210467685826e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.51303540853224e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00700310000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.16359780371051e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00700310000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.06972436836688e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.44902184277485e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00458430000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.14026918120179e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00458430000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.05751762048023e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.15354919573938e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.27171044429298e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.04596258367745e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.48482629671852e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00608530000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.04101793714555e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00608530000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.05471947709634e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.10748582920366e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00540865000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.14463447968037e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00540865000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.03478780531702e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.30022504049707e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 4.14938496401267e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.02356911397719e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.91383952018231e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00803525000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.93968063427158e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00803525000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.03962913672458e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.60046720275456e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.98015611165795e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.02112161195874e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.67182768989269e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.91222805575044e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.01921097566349e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.83454125158810e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.89664409112637e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.00966765663611e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.77711525796947e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.88279624773343e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.00873669242911e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.04578566273172e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.84630190333160e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.00338964968508e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.24286780337865e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.65270882738885e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -1.00548270999974e+000;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.85311551950812e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00602790000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.85361260541871e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00602790000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.81064380661485e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.54424197074739e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.78523848519499e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.86732199629356e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.86438771440628e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00597370000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.68933143295441e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00597370000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.84778030180111e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.62949110771949e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.53730002742149e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.92008607902471e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.93108186682517e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.66812077341225e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.91524535060015e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.39677234686427e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.55653858414935e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.76386237608478e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.78483520572257e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.61878570011739e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.60182128618071e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.95281591968541e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00704710000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.45853316084792e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00704710000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.59943801813146e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.04740357264544e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00585580000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.56922271439770e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00585580000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.61492535973584e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.28460076454382e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.31190356500382e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.46996081910597e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.69252086931227e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.27946612517449e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.55480984850200e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.67866265761320e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00579840000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.31984240471660e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00579840000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.40050252577566e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.08377716750347e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00447965000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.30315048099844e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00447965000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.41464144982751e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.49716340667385e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00795000000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.33257992109106e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00795000000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.43306954986025e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.53864392052222e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.31033501119999e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.37533074838942e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.29164263500278e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00469570000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.24629176092739e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00469570000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.25943628928205e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.94331488220629e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.28217356576815e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.21775199879567e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.45117913290357e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00469570000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.36916383183443e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00469570000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.17197515235016e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.86521740787245e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00706030000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.23937839640877e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00706030000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.23445848859526e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.16359458485596e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.20640754164600e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.20276305538583e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.88912275376979e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00388295000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.08591865292047e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00388295000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.14200579291614e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.11496218661243e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.16439073480877e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.14335977501310e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.26542476304883e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 1.00447945000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.14819321202597e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 1.00447945000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.96677179970495e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.76089574233334e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00505305000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.96991803642241e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00505305000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.02164425625018e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.79478231167505e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00919210000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.13589556647171e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00919210000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -9.09090420915857e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.04513592069990e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00338980000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.03748201557120e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00338980000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.84863183097099e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.51658446507535e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00450575000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.98626283314928e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00450575000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.93788877746115e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.25633053428663e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.00663554034454e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.97092132194790e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.98153145484898e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00338985000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 3.05074637526339e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00338985000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.80711863826963e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.82987941874910e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.92538677438954e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.89844814737609e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.87530988991696e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00574120000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.87667136755009e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00574120000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.88344337982879e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.00683494573026e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00932010000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.76165621823566e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00932010000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.98955755932498e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -9.70094924248699e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.83326696253928e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.81842519711517e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.50888140446814e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.77181096788150e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.74456769305413e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.06906922116498e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.70393495396690e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.85220397938889e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.43521014331917e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.66583349042253e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.77125137574799e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.88770030390538e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00687670000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.78363436814019e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00687670000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.74886152219520e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.04070892923905e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.63011611394581e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.70159791203798e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.59409294308114e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.57261593466911e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.62907708930122e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.18180085505319e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.66927456258883e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.65871941550981e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.33357618289883e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00447965000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.53203840719392e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00447965000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.56058473087825e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.35505775208303e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.62387366317930e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.00597380000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.64055483382314e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.77470517054178e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 1.00670825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.61927921667936e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 1.00670825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.59033069944917e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.90787070242997e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00447935000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.52399963252161e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00447935000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.45463514851713e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.21073890881942e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00450575000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.48137359418328e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00450575000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.61642260136771e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.87236721400883e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00540850000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.47497607489960e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00540850000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.51752598320242e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.18044074947461e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 1.00783855000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.51959400520111e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 1.00783855000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.55381410356535e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.67552620974728e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.38241148447355e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.55778233421664e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.13062431022638e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.55667927643158e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.43562937968402e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.77191409599249e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00723245000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.37733827780304e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00723245000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.48584677435336e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -8.77187715769809e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00670825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.39759939051671e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00670825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.43354909393083e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.79767188970298e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.01021455000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.33119624680500e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.01021455000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.52335396554461e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.22915716580218e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00786315000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.34228916370501e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00786315000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.41871134082734e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.96466475559981e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00453715000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.24356544802321e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00453715000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.39925765472793e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.74003980645018e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.29441329673712e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.36369085659304e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.86996091792572e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.28784104524686e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.40597680105433e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.44155944335898e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.26889596671220e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.34281395913340e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.76417011888033e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00579855000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.28411025058269e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00579855000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.35404067331162e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.52844926073597e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00585595000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.24713391608532e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00585595000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.27250961246524e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.18195903172522e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.31418178832382e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.30038747665049e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.29524780730207e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00579855000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.24293284177136e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00579855000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.25767901052030e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.48138876030253e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.25021524453584e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.31079444845892e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.91848509796843e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.16392053701541e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.23067195966086e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.73304048932484e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.08426420300036e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.20985204569035e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.16623444432321e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00874170000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.20621625763134e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00874170000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.21120926730387e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.81672925570510e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00574100000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.08152292934228e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00574100000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.15294491025172e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.83925754246277e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.15029664479752e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.21496637148657e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.28533736400551e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.02106054729883e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.24720861150105e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.55190542658536e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.15196975310379e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.18963876390658e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.81528457815337e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00388300000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.15450773565300e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00388300000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.03910317844790e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.98346440408712e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00320465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.19526092630696e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00320465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.01664135410525e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.48788905831982e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00298825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.11815977008700e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00298825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.95181771860899e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.94955137200370e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 9.35605050000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.55700112146255e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 9.35605050000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.23873863921763e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.44731701914128e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597370000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.08878109388303e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597370000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.01472079380837e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.87628768695591e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.29895850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.90426600285277e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.29895850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.25296211672365e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.32057095704809e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.33512750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.71180854180527e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.33512750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.26592693611361e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * 8.53940513455344e-004;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 9.34997300000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.61413566503211e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 9.34997300000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.22177188729351e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.00476930363413e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.09120376055360e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.07383367738933e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.86916316409543e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 9.36917650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.44343541445297e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 9.36917650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.18247019701182e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * 6.66050501052530e-004;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 1.01083425000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.02139306226826e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 1.01083425000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.15038385501572e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.56998286389470e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.08105107275693e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00677365000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.98433312331508e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.83332710960872e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 9.33343600000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.69254794358145e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 9.33343600000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.21297201561670e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.61223917905374e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.46230000000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -6.45415514949433e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.46230000000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.43087314272682e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.47772644889763e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.08306327781507e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00710475000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.09229087292379e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.26083312710763e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00388300000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.12225149309302e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00388300000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.97722380029218e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.56643914360882e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.35265750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.51917464662619e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.35265750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.32063288689363e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.17716317173935e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00453715000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.15300619053215e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00453715000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.00450564063793e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.47394570659403e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 9.33361650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.65888771352939e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 9.33361650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.32350342579400e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.33659217687485e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00338985000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.08861759498126e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00338985000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.91534225252993e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.01684445565853e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.05885135831362e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.05539928921007e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.54471023169773e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.29782850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.90053791826895e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.29782850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.18515879608730e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.42000946910426e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.33534000000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.59867089771222e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.33534000000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.13657474594585e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.25267451673284e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.03450271708349e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00580515000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.96076759298619e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.61099978757922e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 9.43685350000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -6.65704184126238e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 9.43685350000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.43790105968432e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.42982021736796e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00320465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.05992927157264e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00320465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.95518857132009e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.78100044291655e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.06581214494166e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00931045000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.03959804087809e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.44903987223049e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 9.33533800000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.44940932969372e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 9.33533800000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.22796870307265e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.09556012242303e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.33343650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.55617612499911e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.33343650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.15572582369647e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.13279767983697e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.91028492234482e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00716015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.07367895692501e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * 1.64918993084266e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.33935250000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.46875063071318e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.33935250000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.17732262503051e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.21230466004240e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 1.00450575000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.11792238787654e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 1.00450575000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.97032040565941e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.59375872691450e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.01021455000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.05247886172917e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.01021455000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.08155103975007e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.17798483428959e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.33992650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.40526715825598e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.33992650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.22853753867294e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.15218309665227e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00356170000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.91192267335049e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00356170000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.84686188927383e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.16688216384171e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00388300000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.96540562578749e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00388300000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.88297894917245e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.88753454065668e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00453695000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.13182837703337e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00453695000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.85350314438359e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.55234996179196e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.00807371307926e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.86204713388704e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.94692026693736e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.29895850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.94958041811786e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.29895850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.14429418178390e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * 3.01195066477527e-003;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 9.29160250000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.75926823056205e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 9.29160250000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.13958936861502e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.10875789219209e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00783855000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.91522133771213e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00783855000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -8.01155107476313e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -9.50250042160736e-003;
//
//    }
//    if ( _y9__ != NULL && *_y9__ <= 4.99544500000000e-001 ) {
//        _PredictProb[2]  += _LearningRate * -9.41356486770220e-001;
//
//    }
//    else if ( _y9__ != NULL && *_y9__ > 4.99544500000000e-001 ) {
//        _PredictProb[2]  += _LearningRate * 1.54843406230415e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.73000301258821e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.34997200000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.35316051432138e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.34997200000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.13104745505093e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.33816931834820e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00469570000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.05793259382363e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00469570000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.84840100943595e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * 5.32663807657898e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00338985000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.09355812572592e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00338985000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.69588975784527e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.38106452486090e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.46208750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -6.23608040884540e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.46208750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.34156892634981e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.20225345521057e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00447945000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.97692501617508e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00447945000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.89663354676044e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.68489053410847e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.33512750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.25163398606178e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.33512750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.11183278598620e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.73477810522536e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.43573800000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -6.55009424510688e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.43573800000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.25114945965734e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.29340252559438e-002;
//
//    }
//    if ( _z7__ != NULL && *_z7__ <= 9.29217550000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.60278594441726e-001;
//
//    }
//    else if ( _z7__ != NULL && *_z7__ > 9.29217550000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.04441561169736e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.21552332651699e-002;
//
//    }
//    if ( _z5__ != NULL && *_z5__ <= 9.37751700000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.07890022358376e-001;
//
//    }
//    else if ( _z5__ != NULL && *_z5__ > 9.37751700000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.15138115661424e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -8.95511634939244e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00490060000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.96595465196858e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00490060000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.75273272934487e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.41154846335950e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00670825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.96122427837413e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00670825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.89621928649464e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * 4.63440653894511e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00201835000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.93555022369446e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00201835000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.72839308231500e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.89456361260034e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00308015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.92677364721342e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00308015000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.74075976019875e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.38223749601638e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 9.42330350000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -6.59507466777379e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 9.42330350000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.27319364053855e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.74680054861409e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.33935250000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.50429812255294e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.33935250000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.00787157506678e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.97649660627742e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.83606503812467e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00723235000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.87100657721439e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.92039334432319e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00786315000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.91335537545045e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00786315000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.92540974510112e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -4.32665751956742e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 9.45383000000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -6.32555960661409e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 9.45383000000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.36996003090166e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -6.34149490144672e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.46208850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -6.22102885765310e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.46208850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.31041660637489e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.46550168873154e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.33343700000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.39661551524839e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.33343700000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.10605474083723e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.48288466140983e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00919210000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.97621089615795e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00919210000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.86717216321649e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.57806597551939e-002;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= 4.86926500000000e-001 ) {
//        _PredictProb[2]  += _LearningRate * -9.55417197179526e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > 4.86926500000000e-001 ) {
//        _PredictProb[2]  += _LearningRate * 1.50300467455585e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.60322660824031e-002;
//
//    }
//    if ( _z2__ != NULL && *_z2__ <= 9.29091600000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.70045446508848e-001;
//
//    }
//    else if ( _z2__ != NULL && *_z2__ > 9.29091600000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 1.97220697614600e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.51881823026917e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.29895700000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.71160078461374e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.29895700000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 1.91769567722702e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.95723900509948e-002;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.01084570000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.83515976206021e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.01084570000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.94925534115510e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.92876879570350e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.29895850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.68211055974907e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.29895850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.12642615881805e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.08632983739049e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.96569227756440e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00597390000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.82443695844852e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.22436171553464e-003;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00574120000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.96021507364749e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00574120000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.82925784706854e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.48585060757337e-002;
//
//    }
//    if ( _z6__ != NULL && *_z6__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 2.03004110693032e-001;
//
//    }
//    else if ( _z6__ != NULL && *_z6__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.69778467739546e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.10619021170483e-002;
//
//    }
//    if ( _z8__ != NULL && *_z8__ <= 9.33992650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.22994064359721e-001;
//
//    }
//    else if ( _z8__ != NULL && *_z8__ > 9.33992650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 1.98936875097601e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.33023886319031e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 9.43972100000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -6.27439770997584e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 9.43972100000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.32900398324897e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.05043072323069e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00704710000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.93438113584635e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00704710000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.79714133129173e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.07192357844408e-003;
//
//    }
//    if ( _y10__ != NULL && *_y10__ <= 4.86648000000000e-001 ) {
//        _PredictProb[2]  += _LearningRate * -9.58328490658574e-001;
//
//    }
//    else if ( _y10__ != NULL && *_y10__ > 4.86648000000000e-001 ) {
//        _PredictProb[2]  += _LearningRate * 1.56784912219578e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * 6.75758538659025e-003;
//
//    }
//    if ( _y__ != NULL && *_y__ <= 4.82699000000000e-001 ) {
//        _PredictProb[2]  += _LearningRate * -9.41721931403688e-001;
//
//    }
//    else if ( _y__ != NULL && *_y__ > 4.82699000000000e-001 ) {
//        _PredictProb[2]  += _LearningRate * 1.54008317315858e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.46155213940846e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 9.29160200000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.60874575032087e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 9.29160200000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 1.95224217638123e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.49755023294457e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00388300000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.91547700816666e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00388300000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.61172599610333e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.93635909912204e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.48388050000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -5.85351537385754e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.48388050000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.30873953451555e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.88632016589304e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.01118970000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.93870996503918e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.01118970000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.84043561769768e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.48185419176038e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.48547200000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -5.89673676184119e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.48547200000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.33798917545174e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.05283476397947e-002;
//
//    }
//    if ( _z3__ != NULL && *_z3__ <= 1.00473205000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.87587516779037e-001;
//
//    }
//    else if ( _z3__ != NULL && *_z3__ > 1.00473205000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.65857818254779e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -5.56698795866999e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.33343650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.15457531904104e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.33343650000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.05627709898373e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -9.94605704654850e-003;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 1.00670825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.81710216902800e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 1.00670825000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.81518705773248e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.16351288472191e-002;
//
//    }
//    if ( _z4__ != NULL && *_z4__ <= 9.35208600000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.11128604428773e-001;
//
//    }
//    else if ( _z4__ != NULL && *_z4__ > 9.35208600000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 1.96841006439878e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.42928719744757e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.46208750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -5.91770837972810e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.46208750000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 2.23583124311205e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.01462845950815e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.94896172999410e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00710470000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.79258295123629e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.52270700057354e-002;
//
//    }
//    if ( _z__ != NULL && *_z__ <= 9.29232550000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.58323252770620e-001;
//
//    }
//    else if ( _z__ != NULL && *_z__ > 9.29232550000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 1.95569875847754e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -3.22078761075364e-002;
//
//    }
//    if ( _z1__ != NULL && *_z1__ <= 9.29895850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * -7.50064713173801e-001;
//
//    }
//    else if ( _z1__ != NULL && *_z1__ > 9.29895850000000e+000 ) {
//        _PredictProb[2]  += _LearningRate * 1.95617401024550e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -7.75812602598367e-003;
//
//    }
//    if ( _z9__ != NULL && *_z9__ <= 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.97425920621857e-001;
//
//    }
//    else if ( _z9__ != NULL && *_z9__ > 1.00467465000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.75220564981524e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -2.56169637756051e-002;
//
//    }
//    if ( _z10__ != NULL && *_z10__ <= 1.00932010000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * 1.86752658510674e-001;
//
//    }
//    else if ( _z10__ != NULL && *_z10__ > 1.00932010000000e+001 ) {
//        _PredictProb[2]  += _LearningRate * -7.80672545483156e-001;
//
//    }
//    else {
//    _PredictProb[2]  += _LearningRate * -1.75480292232479e-002;
//
//    }
//
//   if(_MaxValue <= _PredictProb[2])
//   {
//     _MaxValue = _PredictProb[2];
//     STRING_SET(_pRet,"Standing");
//   }
//
//}
//
import "C"

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log Is the Default package logger
var log = logger.GetLogger("activity-tibco-inference")

// InferfenceActivity is an Activity that is used to Invoke a ML Model using flogo-ml framework
type ModelActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a New InferfenceActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ModelActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *ModelActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Runs an ML model
func (a *ModelActivity) Eval(context activity.Context) (done bool, err error) {

	var i1 = C.double(float64(context.GetInput("_x__").(float64)))
	var i2 = C.double(float64(context.GetInput("_y__").(float64)))
	var i3 = C.double(float64(context.GetInput("_z__").(float64)))
	var i4 = C.double(float64(context.GetInput("_x1__").(float64)))
	var i5 = C.double(float64(context.GetInput("_x2__").(float64)))
	var i6 = C.double(float64(context.GetInput("_x3__").(float64)))
	var i7 = C.double(float64(context.GetInput("_x4__").(float64)))
	var i8 = C.double(float64(context.GetInput("_x5__").(float64)))
	var i9 = C.double(float64(context.GetInput("_x6__").(float64)))
	var i10 = C.double(float64(context.GetInput("_x7__").(float64)))
	var i11 = C.double(float64(context.GetInput("_x8__").(float64)))
	var i12 = C.double(float64(context.GetInput("_x9__").(float64)))
	var i13 = C.double(float64(context.GetInput("_x10__").(float64)))
	var i14 = C.double(float64(context.GetInput("_y1__").(float64)))
	var i15 = C.double(float64(context.GetInput("_y2__").(float64)))
	var i16 = C.double(float64(context.GetInput("_y3__").(float64)))
	var i17 = C.double(float64(context.GetInput("_y4__").(float64)))
	var i18 = C.double(float64(context.GetInput("_y5__").(float64)))
	var i19 = C.double(float64(context.GetInput("_y6__").(float64)))
	var i20 = C.double(float64(context.GetInput("_y7__").(float64)))
	var i21 = C.double(float64(context.GetInput("_y8__").(float64)))
	var i22 = C.double(float64(context.GetInput("_y9__").(float64)))
	var i23 = C.double(float64(context.GetInput("_y10__").(float64)))
	var i24 = C.double(float64(context.GetInput("_z1__").(float64)))
	var i25 = C.double(float64(context.GetInput("_z2__").(float64)))
	var i26 = C.double(float64(context.GetInput("_z3__").(float64)))
	var i27 = C.double(float64(context.GetInput("_z4__").(float64)))
	var i28 = C.double(float64(context.GetInput("_z5__").(float64)))
	var i29 = C.double(float64(context.GetInput("_z6__").(float64)))
	var i30 = C.double(float64(context.GetInput("_z7__").(float64)))
	var i31 = C.double(float64(context.GetInput("_z8__").(float64)))
	var i32 = C.double(float64(context.GetInput("_z9__").(float64)))
	var i33 = C.double(float64(context.GetInput("_z10__").(float64)))

	var result C.char
	C._BTrees_Prediction_T14_16_28(&i1, &i2, &i3, &i4, &i5, &i6, &i7, &i8, &i9, &i10, &i11, &i12, &i13, &i14, &i15, &i16, &i17, &i18, &i19, &i20, &i21, &i22, &i23, &i24, &i25, &i26, &i27, &i28, &i29, &i30, &i31, &i32, &i33, &result)

	context.SetOutput("result", C.GoString(&result))

	return true, nil
}

